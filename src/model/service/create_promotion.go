package service

import (
	"strings"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *promotionDomainService) CreatePromotion(promotionDomain model.PromotionDomainInterface) *rest_err.RestErr {
	promotionDomain.SetCode(strings.ToUpper(promotionDomain.GetCode()))

	return pd.repository.CreatePromotion(promotionDomain)
}

func boolPointer(value bool) *bool {
	return &value
}
