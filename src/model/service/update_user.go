package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	if userDomain.GetPassword() != "" {
		if err := userDomain.EncryptPassword(); err != nil {
			return nil, err
		}
	}

	return ud.repository.UpdateUser(userId, userDomain)
}
