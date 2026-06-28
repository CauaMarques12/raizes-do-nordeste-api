package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewProductControllerInterface(
	serviceInterface service.ProductDomainService,
) ProductControllerInterface {
	return &productControllerInterface{
		service: serviceInterface,
	}
}

type ProductControllerInterface interface {
	CreateProduct(c *gin.Context)
	FindProductById(c *gin.Context)
	FindProducts(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

type productControllerInterface struct {
	service service.ProductDomainService
}
