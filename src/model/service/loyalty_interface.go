package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewLoyaltyDomainService() LoyaltyDomainService {
	return &loyaltyDomainService{
		loyaltyRepository: repository.NewLoyaltyRepository(),
		userRepository:    repository.NewUserRepository(),
	}
}

type loyaltyDomainService struct {
	loyaltyRepository repository.LoyaltyRepository
	userRepository    repository.UserRepository
}

type LoyaltyDomainService interface {
	FindBalance(string) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr)
	FindMovements(string) ([]model.LoyaltyMovementDomainInterface, *rest_err.RestErr)
	RedeemPoints(string, int64) (model.LoyaltyMovementDomainInterface, *rest_err.RestErr)
}
