package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewStockControllerInterface(
	serviceInterface service.StockDomainService,
) StockControllerInterface {
	return &stockControllerInterface{
		service: serviceInterface,
	}
}

type StockControllerInterface interface {
	CreateStockMovement(c *gin.Context)
	FindStockBalance(c *gin.Context)
}

type stockControllerInterface struct {
	service service.StockDomainService
}
