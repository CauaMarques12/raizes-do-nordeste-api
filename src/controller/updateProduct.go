package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (pc *productControllerInterface) UpdateProduct(c *gin.Context) {
	productID := c.Param("productId")
	var productRequest request.ProductUpdateRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewProductUpdateDomain(
		productRequest.Name,
		productRequest.Description,
		productRequest.Category,
		productRequest.PriceCents,
	)

	updatedProduct, err := pc.service.UpdateProduct(productID, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(updatedProduct))
}
