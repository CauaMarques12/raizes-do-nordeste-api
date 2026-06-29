package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model/repository"
)

func NewUnitDomainService() UnitDomainService {
	return &unitDomainService{
		repository: repository.NewUnitRepository(),
	}
}

type unitDomainService struct {
	repository repository.UnitRepository
}

type UnitDomainService interface {
	CreateUnit(model.UnitDomainInterface) *rest_err.RestErr
	FindUnit(string) (model.UnitDomainInterface, *rest_err.RestErr)
	FindUnits(int64, int64) ([]model.UnitDomainInterface, *rest_err.RestErr)
	UpdateUnit(string, model.UnitDomainInterface) (model.UnitDomainInterface, *rest_err.RestErr)
}
