package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (uc *unitControllerInterface) CreateUnit(c *gin.Context) {
	var unitRequest request.UnitRequest
	if err := c.ShouldBindJSON(&unitRequest); err != nil {
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
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertUnitDomainToResponse(domain))
}
