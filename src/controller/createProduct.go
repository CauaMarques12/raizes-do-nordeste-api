package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (pc *productControllerInterface) CreateProduct(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewProductDomain(
		productRequest.Name,
		productRequest.Description,
		productRequest.Category,
		productRequest.PriceCents,
	)

	if err := pc.service.CreateProduct(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertProductDomainToResponse(domain))
}
