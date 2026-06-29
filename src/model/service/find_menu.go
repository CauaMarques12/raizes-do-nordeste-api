package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (md *menuDomainService) FindMenuByUnit(unitID, category string, page, limit int64) ([]model.MenuItemDomainInterface, *rest_err.RestErr) {
	if _, err := md.unitRepository.FindUnitByID(unitID); err != nil {
		return nil, err
	}

	products, err := md.productRepository.FindProducts(category, page, limit)
	if err != nil {
		return nil, err
	}

	menuItems := make([]model.MenuItemDomainInterface, 0, len(products))
	for _, product := range products {
		quantity := int64(0)
		balance, err := md.stockRepository.FindBalance(unitID, product.GetID())
		if err != nil {
			if err.Code != 404 {
				return nil, err
			}
		} else {
			quantity = balance.GetQuantity()
		}

		menuItems = append(menuItems, model.NewMenuItemDomain(
			product.GetID(),
			product.GetName(),
			product.GetDescription(),
			product.GetCategory(),
			product.GetPriceCents(),
			quantity,
		))
	}

	return menuItems, nil
}
