package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *promotionDomainService) FindPromotion(promotionID string) (model.PromotionDomainInterface, *rest_err.RestErr) {
	return pd.repository.FindPromotionByID(promotionID)
}

func (pd *promotionDomainService) FindPromotions(active *bool, page, limit int64) ([]model.PromotionDomainInterface, *rest_err.RestErr) {
	return pd.repository.FindPromotions(active, page, limit)
}
