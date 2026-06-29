package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *stockControllerInterface) FindStockBalance(c *gin.Context) {
	logger.Info("Iniciando consulta de estoque", zap.String("jornada", "find_stock"))
	unitID := c.Query("unidadeId")
	productID := c.Query("produtoId")
	if unitID == "" || productID == "" {
		err := rest_err.NewBadRequestError("unidadeId and produtoId are required")
		c.JSON(err.Code, err)
		return
	}

	stockDomain, err := sc.service.FindBalance(unitID, productID)
	if err != nil {
		logger.Error("Erro ao tentar consultar estoque", err, zap.String("jornada", "find_stock"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertStockBalanceDomainToResponse(stockDomain))
}
