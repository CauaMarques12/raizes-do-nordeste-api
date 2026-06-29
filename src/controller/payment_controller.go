package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewPaymentControllerInterface(
	serviceInterface service.PaymentDomainService,
) PaymentControllerInterface {
	return &paymentControllerInterface{
		service: serviceInterface,
	}
}

type PaymentControllerInterface interface {
	CreatePayment(c *gin.Context)
	FindPaymentById(c *gin.Context)
	FindPaymentsByOrderId(c *gin.Context)
}

type paymentControllerInterface struct {
	service service.PaymentDomainService
}
