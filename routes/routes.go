package routes

import (
	"cota_frete/config"
	"cota_frete/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database, cfg *config.Config) {
	r.POST("/quote", func(c *gin.Context) {
		handlers.QuoteHandler(c, db, cfg)
	})

	r.GET("/metrics", func(c *gin.Context) {
		handlers.GetMetricsHandler(c, db)
	})
}
