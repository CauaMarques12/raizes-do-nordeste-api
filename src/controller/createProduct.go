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

func (pc *productControllerInterface) CreateProduct(c *gin.Context) {
	logger.Info("Iniciando controlador de criacao de produto", zap.String("jornada", "create_product"))
	var productRequest request.ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		logger.Error("Erro ao tentar validar as informacoes do produto", err, zap.String("jornada", "create_product"))
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
		logger.Error("Erro ao tentar criar produto", err, zap.String("jornada", "create_product"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Produto criado com sucesso", zap.String("jornada", "create_product"))
	c.JSON(http.StatusCreated, view.ConvertProductDomainToResponse(domain))
}
