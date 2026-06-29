package service

import (
	"fmt"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (od *orderDomainService) CreateOrder(orderDomain model.OrderDomainInterface) *rest_err.RestErr {
	if _, err := od.userRepository.FindUserByID(orderDomain.GetClientID()); err != nil {
		return err
	}
	if _, err := od.unitRepository.FindUnitByID(orderDomain.GetUnitID()); err != nil {
		return err
	}

	items, totalCents, requiredByProduct, err := od.prepareOrderItems(orderDomain.GetItems())
	if err != nil {
		return err
	}

	if err := od.validateStock(orderDomain.GetUnitID(), requiredByProduct); err != nil {
		return err
	}

	if err := od.decreaseStock(orderDomain.GetUnitID(), requiredByProduct); err != nil {
		return err
	}

	orderDomain.SetItems(items)
	orderDomain.SetTotalCents(totalCents)

	return od.orderRepository.CreateOrder(orderDomain)
}

func (od *orderDomainService) prepareOrderItems(
	items []model.OrderItemDomain,
) ([]model.OrderItemDomain, int64, map[string]int64, *rest_err.RestErr) {
	preparedItems := make([]model.OrderItemDomain, 0, len(items))
	requiredByProduct := map[string]int64{}
	var totalCents int64

	for _, item := range items {
		product, err := od.productRepository.FindProductByID(item.ProductID)
		if err != nil {
			return nil, 0, nil, err
		}

		subtotal := product.GetPriceCents() * item.Quantity
		preparedItems = append(preparedItems, model.OrderItemDomain{
			ProductID:      item.ProductID,
			Quantity:       item.Quantity,
			UnitPriceCents: product.GetPriceCents(),
			SubtotalCents:  subtotal,
		})
		requiredByProduct[item.ProductID] += item.Quantity
		totalCents += subtotal
	}

	return preparedItems, totalCents, requiredByProduct, nil
}

func (od *orderDomainService) validateStock(unitID string, requiredByProduct map[string]int64) *rest_err.RestErr {
	for productID, requiredQuantity := range requiredByProduct {
		balance, err := od.stockRepository.FindBalance(unitID, productID)
		if err != nil {
			if err.Code == 404 {
				return rest_err.NewConflictError("Stock is not available")
			}
			return err
		}
		if balance.GetQuantity() < requiredQuantity {
			return rest_err.NewConflictError("Insufficient stock")
		}
	}

	return nil
}

func (od *orderDomainService) decreaseStock(unitID string, requiredByProduct map[string]int64) *rest_err.RestErr {
	for productID, requiredQuantity := range requiredByProduct {
		movement := model.NewStockMovementDomain(
			unitID,
			productID,
			"SAIDA",
			requiredQuantity,
			fmt.Sprintf("Criacao de pedido"),
		)
		if err := od.stockRepository.CreateMovement(movement); err != nil {
			return err
		}
	}

	return nil
}
