package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewAuthDomainService() AuthDomainService {
	return &authDomainService{
		repository: repository.NewUserRepository(),
	}
}

type authDomainService struct {
	repository repository.UserRepository
}

type AuthDomainService interface {
	Login(email, password string) (string, int64, model.UserDomainInterface, *rest_err.RestErr)
}
