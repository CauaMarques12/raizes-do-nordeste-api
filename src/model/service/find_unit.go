package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ud *unitDomainService) FindUnit(unitID string) (model.UnitDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUnitByID(unitID)
}

func (ud *unitDomainService) FindUnits(page, limit int64) ([]model.UnitDomainInterface, *rest_err.RestErr) {
	return ud.repository.FindUnits(page, limit)
}
