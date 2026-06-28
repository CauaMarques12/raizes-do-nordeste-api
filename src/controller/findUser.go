package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindLoggedUser(c *gin.Context) {
	userID := c.GetString("userId")
	userDomain, err := uc.service.FindUser(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userID := c.Param("userId")
	userDomain, err := uc.service.FindUser(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		err := rest_err.NewBadRequestError("Email is required")
		c.JSON(err.Code, err)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
