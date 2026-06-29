package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (sd *stockDomainService) CreateMovement(
	stockMovementDomain model.StockMovementDomainInterface,
) *rest_err.RestErr {
	if _, err := sd.unitRepository.FindUnitByID(stockMovementDomain.GetUnitID()); err != nil {
		return err
	}
	if _, err := sd.productRepository.FindProductByID(stockMovementDomain.GetProductID()); err != nil {
		return err
	}

	return sd.stockRepository.CreateMovement(stockMovementDomain)
}
