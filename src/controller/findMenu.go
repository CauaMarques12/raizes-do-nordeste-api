package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (mc *menuControllerInterface) FindMenuByUnit(c *gin.Context) {
	logger.Info("Iniciando consulta de cardapio por unidade", zap.String("jornada", "find_menu"))
	unitID := c.Param("unitId")
	category := c.Query("category")
	page, limit, err := getPagination(c)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	menuItems, restErr := mc.service.FindMenuByUnit(unitID, category, page, limit)
	if restErr != nil {
		logger.Error("Erro ao tentar consultar cardapio por unidade", restErr, zap.String("jornada", "find_menu"))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, view.ConvertMenuItemDomainsToResponse(menuItems))
}
