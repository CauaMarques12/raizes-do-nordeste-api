package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (*userDomainService) FindUser(string) (
	 model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
