package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (od *orderDomainService) UpdateStatus(orderID, status string) (model.OrderDomainInterface, *rest_err.RestErr) {
	return od.orderRepository.UpdateStatus(orderID, status)
}

func (od *orderDomainService) CancelOrder(orderID string) (model.OrderDomainInterface, *rest_err.RestErr) {
	return od.orderRepository.UpdateStatus(orderID, "CANCELADO")
}
