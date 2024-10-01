package main

import (
	"cota_frete/config"
	"cota_frete/db"
	"cota_frete/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database := db.ConnectDatabase(cfg)
	defer db.DisconnectDatabase(database.Client())
	r := gin.Default()

	routes.RegisterRoutes(r, database, cfg)

	log.Println(fmt.Sprintf("Server iniciado na porta :%s", cfg.ApiPort))
	log.Fatal(r.Run(":" + cfg.ApiPort))
}
