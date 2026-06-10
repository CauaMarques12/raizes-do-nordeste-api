package service

import (
	"fmt"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/logger"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.uber.org/zap"
)



func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Iniciando create user model", zap.String("jornada", "create_user"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
