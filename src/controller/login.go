package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/response"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ac *authControllerInterface) Login(c *gin.Context) {
	logger.Info("Iniciando controlador de login", zap.String("jornada", "login"))
	var authRequest request.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		logger.Error("Erro ao tentar validar as informacoes de login", err, zap.String("jornada", "login"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	token, expiresIn, userDomain, err := ac.service.Login(authRequest.Email, authRequest.Password)
	if err != nil {
		logger.Error("Erro ao tentar realizar login", err, zap.String("jornada", "login"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Login realizado com sucesso", zap.String("jornada", "login"))
	c.JSON(http.StatusOK, response.AuthResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
		User:        view.ConvertDomainToResponse(userDomain),
	})
}
