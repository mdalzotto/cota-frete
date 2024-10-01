package config

import (
	"fmt"
	"os"
)

type Config struct {
	ApiPort              string
	MongoHost            string
	MongoPort            string
	MongoDatabase        string
	ApiPath              string
	ApiToken             string
	ApiPlatformCode      string
	ApiRegisteredNumber  string
	ApiDispatcherZipcode string
}

func LoadConfig() *Config {
	return &Config{
		ApiPort:              getEnv("API_PORT", "8080"),
		MongoHost:            getEnv("MONGO_HOST", "mongodb"),
		MongoPort:            getEnv("MONGO_PORT", "27017"),
		MongoDatabase:        getEnv("MONGO_DATABASE", "cota_frete"),
		ApiPath:              getEnv("API_PATH", "https://sp.freterapido.com/api/v3/quote/simulate"),
		ApiToken:             getEnv("API_TOKEN", ""),
		ApiPlatformCode:      getEnv("API_PLATFORM_CODE", ""),
		ApiRegisteredNumber:  getEnv("API_REGISTERED_NUMBER", ""),
		ApiDispatcherZipcode: getEnv("API_DISPATCHER_ZIPCODE", ""),
	}
}

func GetMongoURI(cfg *Config) string {
	return fmt.Sprintf("mongodb://%s:%s/%s", cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
