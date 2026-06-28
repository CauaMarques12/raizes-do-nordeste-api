package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (uc *unitControllerInterface) UpdateUnit(c *gin.Context) {
	unitID := c.Param("unitId")
	var unitRequest request.UnitUpdateRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
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
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertUnitDomainToResponse(updatedUnit))
}
