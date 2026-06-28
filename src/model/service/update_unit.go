package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (ud *unitDomainService) UpdateUnit(
	unitID string,
	unitDomain model.UnitDomainInterface,
) (model.UnitDomainInterface, *rest_err.RestErr) {
	return ud.repository.UpdateUnit(unitID, unitDomain)
}
