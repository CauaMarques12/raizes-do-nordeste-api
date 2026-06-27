package view

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/controller/model/response"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:                   userDomain.GetID(),
		Email:                userDomain.GetEmail(),
		Name:                 userDomain.GetName(),
		Role:                 userDomain.GetRole(),
		FidelidadeConsentida: userDomain.GetFidelidadeConsentida(),
		Active:               userDomain.GetActive(),
		CreatedAt:            userDomain.GetCreatedAt(),
		UpdatedAt:            userDomain.GetUpdatedAt(),
	}
}
