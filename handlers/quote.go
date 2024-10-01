package handlers

import (
	"cota_frete/config"
	"cota_frete/models"
	"cota_frete/repository"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wklken/gorequest"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

func QuoteHandler(c *gin.Context, db *mongo.Database, cfg *config.Config) {
	var req models.CRequest
	if err := bindJSONRequest(c, &req); err != nil {
		return
	}

	zipcode, err := validateZipCode(req.Recipient.Address.ZipCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quoteAPIRequest := buildQuoteAPIRequest(req, cfg, zipcode)
	if err := validateVolumes(c, req.Volumes); err != nil {
		return
	}

	body, err := makeAPIRequest(cfg.ApiPath, quoteAPIRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := processAPIResponse(c, body, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func bindJSONRequest(c *gin.Context, req *models.CRequest) error {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func validateZipCode(zipCode string) (int, error) {
	zip, err := strconv.Atoi(zipCode)
	if err != nil {
		return 0, fmt.Errorf("CEP do recebedor é inválido, o CEP deve conter apenas números")
	}
	return zip, nil
}

func buildQuoteAPIRequest(req models.CRequest, cfg *config.Config, zipcode int) models.QuoteApiRequest {
	quoteAPIRequest := models.QuoteApiRequest{
		Shipper: models.Shipper{
			RegisteredNumber: cfg.ApiRegisteredNumber,
			Token:            cfg.ApiToken,
			PlatformCode:     cfg.ApiPlatformCode,
		},
		Recipient: models.Recipient{
			Type:    0,
			Country: "BRA",
			Zipcode: uint32(zipcode),
		},
		Dispatchers: []models.Dispatcher{
			{
				RegisteredNumber: cfg.ApiRegisteredNumber,
				Zipcode:          29161376,
				Volumes:          make([]models.Volume, len(req.Volumes)),
			},
		},
		SimulationType: []int{0},
	}

	for i, volume := range req.Volumes {
		quoteAPIRequest.Dispatchers[0].Volumes[i] = models.Volume{
			Amount:        volume.Amount,
			Category:      fmt.Sprint(volume.Category),
			UnitaryWeight: volume.UnitaryWeight,
			Price:         volume.Price,
			UnitaryPrice:  float64(volume.Price) / float64(volume.Amount),
			SKU:           volume.SKU,
			Height:        volume.Height,
			Width:         volume.Width,
			Length:        volume.Length,
		}
	}

	return quoteAPIRequest
}

func validateVolumes(c *gin.Context, volumes []models.Volume) error {
	for _, volume := range volumes {
		if volume.Amount <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantidade deve ser maior que zero."})
			return fmt.Errorf("invalid volume amount")
		}
	}
	return nil
}

func makeAPIRequest(apiPath string, quoteAPIRequest models.QuoteApiRequest) (string, error) {
	_, body, errs := gorequest.New().Post(apiPath).Send(quoteAPIRequest).End()
	if len(errs) > 0 {
		return "", errs[0]
	}
	return body, nil
}

func processAPIResponse(c *gin.Context, body string, db *mongo.Database) error {
	var respp models.QuoteAPIResponse
	if err := json.Unmarshal([]byte(body), &respp); err != nil {
		return err
	}

	carrierResponse, hasOffers := extractCarrierDetails(respp)

	if !hasOffers {
		c.JSON(http.StatusOK, gin.H{"message": "Não há ofertas de frete disponíveis no momento."})
		return nil
	}

	quoteResponse := models.QuoteResponse{Response: body}
	err := repository.InsertQuote(c, db, quoteResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gravar dados"})
		return nil
	}

	c.JSON(http.StatusCreated, carrierResponse)
	return nil
}

func extractCarrierDetails(respp models.QuoteAPIResponse) (models.CarrierResponse, bool) {
	var carrierResponse models.CarrierResponse
	hasOffers := false

	for _, dispatcher := range respp.Dispatchers {
		for _, offer := range dispatcher.Offers {
			carrierResponse.Carrier = append(carrierResponse.Carrier, models.CarrierDetail{
				Name:     offer.Carrier.Name,
				Service:  offer.Service,
				Deadline: offer.DeliveryTime.Days,
				Price:    offer.FinalPrice,
			})
			hasOffers = true
		}
	}

	return carrierResponse, hasOffers
}
