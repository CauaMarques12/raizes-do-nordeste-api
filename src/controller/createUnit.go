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

func (uc *unitControllerInterface) CreateUnit(c *gin.Context) {
	logger.Info("Iniciando controlador de criacao de unidade", zap.String("jornada", "create_unit"))
	var unitRequest request.UnitRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
		logger.Error("Erro ao tentar validar as informacoes da unidade", err, zap.String("jornada", "create_unit"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUnitDomain(
		unitRequest.Name,
		unitRequest.Address,
		unitRequest.City,
		unitRequest.State,
	)

	if err := uc.service.CreateUnit(domain); err != nil {
		logger.Error("Erro ao tentar criar unidade", err, zap.String("jornada", "create_unit"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Unidade criada com sucesso", zap.String("jornada", "create_unit"))
	c.JSON(http.StatusCreated, view.ConvertUnitDomainToResponse(domain))
}
