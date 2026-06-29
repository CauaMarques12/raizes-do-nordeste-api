package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ld *loyaltyDomainService) RedeemPoints(userID string, points int64) (model.LoyaltyMovementDomainInterface, *rest_err.RestErr) {
	userDomain, err := ld.userRepository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	if !userDomain.GetFidelidadeConsentida() {
		return nil, rest_err.NewForbiddenError("Usuario nao consentiu participar da fidelidade")
	}

	movement := model.NewLoyaltyMovementDomain(
		userID,
		"RESGATE",
		points,
		"Resgate de pontos de fidelidade",
		"",
	)
	if err := ld.loyaltyRepository.CreateMovement(movement); err != nil {
		return nil, err
	}

	return movement, nil
}
