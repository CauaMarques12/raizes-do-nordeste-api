package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ud *userDomainService) FindUser(userID string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUserByID(userID)
}

func (ud *userDomainService) FindUserByEmail(email string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUserByEmail(email)
}
