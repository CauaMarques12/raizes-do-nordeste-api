package main

import (
	"log"	
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
  "github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
)

func main() {
  logger.Info("Starting Raizes do Nordeste API")
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
 

  router := gin.Default()
  routes.InitRoutes(&router.RouterGroup)
  
  if err := router.Run (":8080"); err != nil {
	log.Fatal(err)
  }

}