package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (od *orderDomainService) FindOrder(orderID string) (model.OrderDomainInterface, *rest_err.RestErr) {
	return od.orderRepository.FindOrderByID(orderID)
}

func (od *orderDomainService) FindOrders(channel, status string, page, limit int64) ([]model.OrderDomainInterface, *rest_err.RestErr) {
	return od.orderRepository.FindOrders(channel, status, page, limit)
}
