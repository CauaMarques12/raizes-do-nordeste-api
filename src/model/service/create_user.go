package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Iniciando create user model", zap.String("jornada", "create_user"))
	if err := userDomain.EncryptPassword(); err != nil {
		return err
	}

	return ud.repository.CreateUser(userDomain)
}
