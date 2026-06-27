package controller

import (
	"net/http"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/validation"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/request"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/view"
	"github.com/gin-gonic/gin"
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
