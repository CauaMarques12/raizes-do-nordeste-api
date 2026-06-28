package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewAuthControllerInterface(
	serviceInterface service.AuthDomainService,
) AuthControllerInterface {
	return &authControllerInterface{
		service: serviceInterface,
	}
}

type AuthControllerInterface interface {
	Login(c *gin.Context)
}

type authControllerInterface struct {
	service service.AuthDomainService
}
