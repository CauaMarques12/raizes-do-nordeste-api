package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (pc *productControllerInterface) FindProductById(c *gin.Context) {
	logger.Info("Iniciando busca de produto por id", zap.String("jornada", "find_product"))
	productID := c.Param("productId")
	productDomain, err := pc.service.FindProduct(productID)
	if err != nil {
		logger.Error("Erro ao tentar buscar produto por id", err, zap.String("jornada", "find_product"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainToResponse(productDomain))
}

func (pc *productControllerInterface) FindProducts(c *gin.Context) {
	logger.Info("Iniciando listagem de produtos", zap.String("jornada", "find_products"))
	category := c.Query("category")
	page, limit, err := getPagination(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	productDomains, err := pc.service.FindProducts(category, page, limit)
	if err != nil {
		logger.Error("Erro ao tentar listar produtos", err, zap.String("jornada", "find_products"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertProductDomainsToResponse(productDomains))
}
