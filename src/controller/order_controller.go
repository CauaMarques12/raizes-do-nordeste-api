package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewOrderControllerInterface(
	serviceInterface service.OrderDomainService,
) OrderControllerInterface {
	return &orderControllerInterface{
		service: serviceInterface,
	}
}

type OrderControllerInterface interface {
	CreateOrder(c *gin.Context)
	FindOrderById(c *gin.Context)
	FindOrders(c *gin.Context)
	UpdateOrderStatus(c *gin.Context)
	CancelOrder(c *gin.Context)
}

type orderControllerInterface struct {
	service service.OrderDomainService
}
