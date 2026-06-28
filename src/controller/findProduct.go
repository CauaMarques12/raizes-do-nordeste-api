package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (pc *productControllerInterface) FindProductById(c *gin.Context) {
	productID := c.Param("productId")
	productDomain, err := pc.service.FindProduct(productID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(productDomain))
}

func (pc *productControllerInterface) FindProducts(c *gin.Context) {
	category := c.Query("category")
	productDomains, err := pc.service.FindProducts(category)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainsToResponse(productDomains))
}
