package service

import (
	"strings"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *promotionDomainService) UpdatePromotion(promotionID string, promotionDomain model.PromotionDomainInterface) (model.PromotionDomainInterface, *rest_err.RestErr) {
	code := promotionDomain.GetCode()
	if code != "" {
		code = strings.ToUpper(code)
	}

	updatedDomain := model.NewPromotionUpdateDomain(
		promotionDomain.GetName(),
		code,
		promotionDomain.GetDiscountPercent(),
		nil,
	)
	if promotionDomain.HasActive() {
		updatedDomain = model.NewPromotionUpdateDomain(
			promotionDomain.GetName(),
			code,
			promotionDomain.GetDiscountPercent(),
			boolPointer(promotionDomain.GetActive()),
		)
	}

	return pd.repository.UpdatePromotion(promotionID, updatedDomain)
}
