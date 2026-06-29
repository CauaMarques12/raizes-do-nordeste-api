package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewMenuControllerInterface(
	serviceInterface service.MenuDomainService,
) MenuControllerInterface {
	return &menuControllerInterface{
		service: serviceInterface,
	}
}

type MenuControllerInterface interface {
	FindMenuByUnit(c *gin.Context)
}

type menuControllerInterface struct {
	service service.MenuDomainService
}
