package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewPromotionDomainService() PromotionDomainService {
	return &promotionDomainService{
		repository: repository.NewPromotionRepository(),
	}
}

type promotionDomainService struct {
	repository repository.PromotionRepository
}

type PromotionDomainService interface {
	CreatePromotion(model.PromotionDomainInterface) *rest_err.RestErr
	FindPromotion(string) (model.PromotionDomainInterface, *rest_err.RestErr)
	FindPromotions(*bool, int64, int64) ([]model.PromotionDomainInterface, *rest_err.RestErr)
	UpdatePromotion(string, model.PromotionDomainInterface) (model.PromotionDomainInterface, *rest_err.RestErr)
}
