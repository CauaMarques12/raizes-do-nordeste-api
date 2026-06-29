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

func (sc *stockControllerInterface) CreateStockMovement(c *gin.Context) {
	logger.Info("Iniciando movimentacao de estoque", zap.String("jornada", "create_stock_movement"))
	var stockRequest request.StockMovementRequest
	if err := c.ShouldBindJSON(&stockRequest); err != nil {
		logger.Error("Erro ao tentar validar movimentacao de estoque", err, zap.String("jornada", "create_stock_movement"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewStockMovementDomain(
		stockRequest.UnitID,
		stockRequest.ProductID,
		stockRequest.Type,
		stockRequest.Quantity,
		stockRequest.Reason,
	)

	if err := sc.service.CreateMovement(domain); err != nil {
		logger.Error("Erro ao tentar movimentar estoque", err, zap.String("jornada", "create_stock_movement"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Movimentacao de estoque realizada com sucesso", zap.String("jornada", "create_stock_movement"))
	c.JSON(http.StatusCreated, view.ConvertStockMovementDomainToResponse(domain))
}
