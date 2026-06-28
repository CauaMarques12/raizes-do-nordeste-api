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

func (uc *unitControllerInterface) UpdateUnit(c *gin.Context) {
	logger.Info("Iniciando controlador de atualizacao de unidade", zap.String("jornada", "update_unit"))
	unitID := c.Param("unitId")
	var unitRequest request.UnitUpdateRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
		logger.Error("Erro ao tentar validar atualizacao de unidade", err, zap.String("jornada", "update_unit"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUnitUpdateDomain(
		unitRequest.Name,
		unitRequest.Address,
		unitRequest.City,
		unitRequest.State,
	)

	updatedUnit, err := uc.service.UpdateUnit(unitID, domain)
	if err != nil {
		logger.Error("Erro ao tentar atualizar unidade", err, zap.String("jornada", "update_unit"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Unidade atualizada com sucesso", zap.String("jornada", "update_unit"))
	c.JSON(http.StatusOK, view.ConvertUnitDomainToResponse(updatedUnit))
}
