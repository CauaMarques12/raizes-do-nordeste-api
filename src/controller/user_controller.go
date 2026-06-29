package controller

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControlleInterface(
	seviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: seviceInterface,
	}
}

type UserControllerInterface interface {
	FindLoggedUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	UpdateLoggedUserFidelityConsent(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
