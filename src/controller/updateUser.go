package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	userID := c.Param("userId")
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Role,
		userRequest.FidelidadeConsentida,
	)

	updatedUser, err := uc.service.UpdateUser(userID, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(updatedUser))
}

func (uc *userControllerInterface) UpdateLoggedUserFidelityConsent(c *gin.Context) {
	logger.Info("Iniciando atualizacao de consentimento de fidelidade", zap.String("jornada", "update_fidelity_consent"))
	userID := c.GetString("userId")
	var consentRequest request.FidelityConsentRequest

	if err := c.ShouldBindJSON(&consentRequest); err != nil {
		logger.Error("Erro ao tentar validar consentimento de fidelidade", err, zap.String("jornada", "update_fidelity_consent"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		"",
		"",
		"",
		"",
		consentRequest.FidelidadeConsentida,
	)

	updatedUser, err := uc.service.UpdateUser(userID, domain)
	if err != nil {
		logger.Error("Erro ao tentar atualizar consentimento de fidelidade", err, zap.String("jornada", "update_fidelity_consent"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Consentimento de fidelidade atualizado com sucesso", zap.String("jornada", "update_fidelity_consent"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(updatedUser))
}
