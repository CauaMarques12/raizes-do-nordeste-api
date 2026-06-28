package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (uc *unitControllerInterface) FindUnitById(c *gin.Context) {
	unitID := c.Param("unitId")
	unitDomain, err := uc.service.FindUnit(unitID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUnitDomainToResponse(unitDomain))
}

func (uc *unitControllerInterface) FindUnits(c *gin.Context) {
	unitDomains, err := uc.service.FindUnits()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUnitDomainsToResponse(unitDomains))
}
