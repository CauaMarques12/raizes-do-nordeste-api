package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ld *loyaltyDomainService) FindBalance(userID string) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr) {
	if _, err := ld.userRepository.FindUserByID(userID); err != nil {
		return nil, err
	}

	balance, err := ld.loyaltyRepository.FindBalance(userID)
	if err != nil {
		if err.Code == 404 {
			return model.NewLoyaltyBalanceDomain(userID, 0), nil
		}
		return nil, err
	}

	return balance, nil
}

func (ld *loyaltyDomainService) FindMovements(userID string) ([]model.LoyaltyMovementDomainInterface, *rest_err.RestErr) {
	if _, err := ld.userRepository.FindUserByID(userID); err != nil {
		return nil, err
	}

	return ld.loyaltyRepository.FindMovements(userID)
}
