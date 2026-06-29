package service

import (
	"fmt"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (od *orderDomainService) UpdateStatus(orderID, status string) (model.OrderDomainInterface, *rest_err.RestErr) {
	orderDomain, err := od.orderRepository.FindOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	if status == "CANCELADO" {
		return od.CancelOrder(orderID)
	}

	if orderDomain.GetStatus() == status {
		return orderDomain, nil
	}

	if !canChangeOrderStatus(orderDomain.GetStatus(), status) {
		return nil, rest_err.NewConflictError("Mudanca de status do pedido nao permitida")
	}

	return od.orderRepository.UpdateStatus(orderID, status)
}

func (od *orderDomainService) CancelOrder(orderID string) (model.OrderDomainInterface, *rest_err.RestErr) {
	orderDomain, err := od.orderRepository.FindOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	if orderDomain.GetStatus() == "CANCELADO" {
		return orderDomain, nil
	}

	if orderDomain.GetStatus() == "ENTREGUE" {
		return nil, rest_err.NewConflictError("Pedido entregue nao pode ser cancelado")
	}

	if err := od.restoreOrderStock(orderDomain); err != nil {
		return nil, err
	}

	return od.orderRepository.UpdateStatus(orderID, "CANCELADO")
}

func canChangeOrderStatus(currentStatus, nextStatus string) bool {
	allowedTransitions := map[string]map[string]bool{
		"AGUARDANDO_PAGAMENTO": {
			"PAGO": true,
		},
		"PAGO": {
			"EM_PREPARO": true,
		},
		"EM_PREPARO": {
			"PRONTO": true,
		},
		"PRONTO": {
			"ENTREGUE": true,
		},
	}

	nextStatuses, exists := allowedTransitions[currentStatus]
	if !exists {
		return false
	}

	return nextStatuses[nextStatus]
}

func (od *orderDomainService) restoreOrderStock(orderDomain model.OrderDomainInterface) *rest_err.RestErr {
	for _, item := range orderDomain.GetItems() {
		movement := model.NewStockMovementDomain(
			orderDomain.GetUnitID(),
			item.ProductID,
			"ENTRADA",
			item.Quantity,
			fmt.Sprintf("Cancelamento do pedido %s", orderDomain.GetID()),
		)
		if err := od.stockRepository.CreateMovement(movement); err != nil {
			return err
		}
	}

	return nil
}
