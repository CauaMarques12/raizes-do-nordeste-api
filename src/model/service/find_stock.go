package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (sd *stockDomainService) FindBalance(
	unitID, productID string,
) (model.StockBalanceDomainInterface, *rest_err.RestErr) {
	if _, err := sd.unitRepository.FindUnitByID(unitID); err != nil {
		return nil, err
	}
	if _, err := sd.productRepository.FindProductByID(productID); err != nil {
		return nil, err
	}

	return sd.stockRepository.FindBalance(unitID, productID)
}
