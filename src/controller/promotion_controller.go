package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewPromotionControllerInterface(
	serviceInterface service.PromotionDomainService,
) PromotionControllerInterface {
	return &promotionControllerInterface{
		service: serviceInterface,
	}
}

type PromotionControllerInterface interface {
	CreatePromotion(c *gin.Context)
	FindPromotionById(c *gin.Context)
	FindPromotions(c *gin.Context)
	UpdatePromotion(c *gin.Context)
}

type promotionControllerInterface struct {
	service service.PromotionDomainService
}
