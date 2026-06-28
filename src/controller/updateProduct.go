package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *productControllerInterface) UpdateProduct(c *gin.Context) {
	logger.Info("Iniciando controlador de atualizacao de produto", zap.String("jornada", "update_product"))
	productID := c.Param("productId")
	var productRequest request.ProductUpdateRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		logger.Error("Erro ao tentar validar atualizacao de produto", err, zap.String("jornada", "update_product"))
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
		logger.Error("Erro ao tentar atualizar produto", err, zap.String("jornada", "update_product"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Produto atualizado com sucesso", zap.String("jornada", "update_product"))
	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(updatedProduct))
}
