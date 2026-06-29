package main

import (
	"log"
	"os"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/database/mongodb"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/routes"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	logger.Info("Starting Raizes do Nordeste API")

	mongodb.InitConnection()

	userService := service.NewUserDomainService()
	authService := service.NewAuthDomainService()
	unitService := service.NewUnitDomainService()
	productService := service.NewProductDomainService()
	stockService := service.NewStockDomainService()
	orderService := service.NewOrderDomainService()
	paymentService := service.NewPaymentDomainService()
	loyaltyService := service.NewLoyaltyDomainService()
	menuService := service.NewMenuDomainService()
	promotionService := service.NewPromotionDomainService()
	userController := controller.NewUserControlleInterface(userService)
	authController := controller.NewAuthControllerInterface(authService)
	unitController := controller.NewUnitControllerInterface(unitService)
	productController := controller.NewProductControllerInterface(productService)
	stockController := controller.NewStockControllerInterface(stockService)
	orderController := controller.NewOrderControllerInterface(orderService)
	paymentController := controller.NewPaymentControllerInterface(paymentService)
	loyaltyController := controller.NewLoyaltyControllerInterface(loyaltyService)
	menuController := controller.NewMenuControllerInterface(menuService)
	promotionController := controller.NewPromotionControllerInterface(promotionService)

	router := gin.Default()
	routes.InitRoutes(
		&router.RouterGroup,
		userController,
		authController,
		unitController,
		productController,
		stockController,
		orderController,
		paymentController,
		loyaltyController,
		menuController,
		promotionController,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
