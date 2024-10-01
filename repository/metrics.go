package repository

import (
	"context"
	"cota_frete/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchQuote(db *mongo.Database, limit int64) ([]models.QuoteResponse, error) {
	collection := db.Collection("quote_responses")
	filter := bson.M{}
	findOptions := options.Find().SetSort(bson.D{{"_id", -1}})
	if limit > 0 {
		findOptions.SetLimit(limit)
	}

	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var quotes []models.QuoteResponse
	if err := cursor.All(context.Background(), &quotes); err != nil {
		return nil, err
	}
	return quotes, nil
}
