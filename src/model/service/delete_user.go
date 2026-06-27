package service

import "github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"

func (ud *userDomainService) DeleteUser(userID string) *rest_err.RestErr {
	return ud.repository.DeleteUser(userID)
}
