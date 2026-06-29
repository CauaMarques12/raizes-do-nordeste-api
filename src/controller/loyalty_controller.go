package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewLoyaltyControllerInterface(
	serviceInterface service.LoyaltyDomainService,
) LoyaltyControllerInterface {
	return &loyaltyControllerInterface{
		service: serviceInterface,
	}
}

type LoyaltyControllerInterface interface {
	FindBalance(c *gin.Context)
	FindHistory(c *gin.Context)
	RedeemPoints(c *gin.Context)
}

type loyaltyControllerInterface struct {
	service service.LoyaltyDomainService
}
