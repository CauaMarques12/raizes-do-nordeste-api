package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *unitControllerInterface) FindUnitById(c *gin.Context) {
	logger.Info("Iniciando busca de unidade por id", zap.String("jornada", "find_unit"))
	unitID := c.Param("unitId")
	unitDomain, err := uc.service.FindUnit(unitID)
	if err != nil {
		logger.Error("Erro ao tentar buscar unidade por id", err, zap.String("jornada", "find_unit"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUnitDomainToResponse(unitDomain))
}

func (uc *unitControllerInterface) FindUnits(c *gin.Context) {
	logger.Info("Iniciando listagem de unidades", zap.String("jornada", "find_units"))
	page, limit, err := getPagination(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	unitDomains, err := uc.service.FindUnits(page, limit)
	if err != nil {
		logger.Error("Erro ao tentar listar unidades", err, zap.String("jornada", "find_units"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUnitDomainsToResponse(unitDomains))
}
