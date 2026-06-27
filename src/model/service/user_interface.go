package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{
		repository: repository.NewUserRepository(),
	}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUser(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
