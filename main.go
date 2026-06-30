package main

import (
	"log"
	"os"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/database/mongodb"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	logger.Info("Starting Raizes do Nordeste API")

	mongodb.InitConnection()

	router := buildRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
