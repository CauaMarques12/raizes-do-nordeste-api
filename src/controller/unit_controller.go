package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUnitControllerInterface(
	serviceInterface service.UnitDomainService,
) UnitControllerInterface {
	return &unitControllerInterface{
		service: serviceInterface,
	}
}

type UnitControllerInterface interface {
	CreateUnit(c *gin.Context)
	FindUnitById(c *gin.Context)
	FindUnits(c *gin.Context)
	UpdateUnit(c *gin.Context)
}

type unitControllerInterface struct {
	service service.UnitDomainService
}
