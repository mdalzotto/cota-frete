package repository

import (
	"cota_frete/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertQuote(c *gin.Context, db *mongo.Database, quote models.QuoteResponse) error {
	collection := db.Collection("quote_responses")
	if _, err := collection.InsertOne(c, quote); err != nil {
		return err
	}
	return nil
}
