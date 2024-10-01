package handlers

import (
	"cota_frete/models"
	"cota_frete/repository"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"math"
	"net/http"
	"strconv"
)

func GetMetricsHandler(c *gin.Context, db *mongo.Database) {
	lastQuotes, err := parseLastQuotes(c.Query("last_quotes"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quotes, err := repository.FetchQuote(db, lastQuotes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados"})
		return
	}

	metricsResponse := calculateMetrics(quotes)

	c.JSON(http.StatusOK, metricsResponse)
}

func parseLastQuotes(lastQuotes string) (int64, error) {
	if lastQuotes == "" {
		return 0, nil
	}
	limit, err := strconv.Atoi(lastQuotes)
	if err != nil || limit <= 0 {
		return 0, fmt.Errorf("Parâmetro last_quotes inválido")
	}
	return int64(limit), nil
}

func calculateMetrics(quotes []models.QuoteResponse) models.OverallMetrics {
	carrierMetrics := make(map[string]models.MetricsResponse)
	cheapestFreight := math.MaxFloat64
	mostExpensiveFreight := 0.0

	for _, quote := range quotes {
		var quoteApiResponse models.QuoteAPIResponse
		if err := json.Unmarshal([]byte(quote.Response), &quoteApiResponse); err != nil {
			continue
		}

		updateCarrierMetrics(quoteApiResponse, carrierMetrics, &cheapestFreight, &mostExpensiveFreight)
	}

	return buildOverallMetrics(carrierMetrics, cheapestFreight, mostExpensiveFreight)
}

func updateCarrierMetrics(quoteApiResponse models.QuoteAPIResponse, carrierMetrics map[string]models.MetricsResponse, cheapestFreight, mostExpensiveFreight *float64) {
	for _, dispatcher := range quoteApiResponse.Dispatchers {
		for _, offer := range dispatcher.Offers {
			carrierName := offer.Carrier.Name

			metric := carrierMetrics[carrierName]
			metric.CarrierName = carrierName
			metric.TotalResults++
			metric.TotalShippingPrice += offer.FinalPrice

			finalPrice := math.Round(offer.FinalPrice*100) / 100

			if finalPrice < *cheapestFreight {
				*cheapestFreight = finalPrice
			}
			if finalPrice > *mostExpensiveFreight {
				*mostExpensiveFreight = finalPrice
			}

			carrierMetrics[carrierName] = metric
		}
	}
}

func buildOverallMetrics(carrierMetrics map[string]models.MetricsResponse, cheapestFreight, mostExpensiveFreight float64) models.OverallMetrics {
	var metricsList []models.MetricsResponse
	for _, metric := range carrierMetrics {
		if metric.TotalResults > 0 {
			metric.AveragePriceShipping = math.Round((metric.TotalShippingPrice/float64(metric.TotalResults))*100) / 100
		}
		metricsList = append(metricsList, metric)
	}

	return models.OverallMetrics{
		TotalCarriers:         len(metricsList),
		CheapestShippingPrice: math.Round(cheapestFreight*100) / 100,
		ShippingMostExpensive: math.Round(mostExpensiveFreight*100) / 100,
		MetricsByCarrier:      metricsList,
	}
}
