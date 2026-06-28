package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/response"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
)

func (ac *authControllerInterface) Login(c *gin.Context) {
	var authRequest request.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	token, expiresIn, userDomain, err := ac.service.Login(authRequest.Email, authRequest.Password)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.AuthResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		User:        view.ConvertDomainToResponse(userDomain),
	})
}
