package service

import (
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *paymentDomainService) FindPayment(paymentID string) (model.PaymentDomainInterface, *rest_err.RestErr) {
	return pd.paymentRepository.FindPaymentByID(paymentID)
}

func (pd *paymentDomainService) FindPaymentsByOrderID(orderID string) ([]model.PaymentDomainInterface, *rest_err.RestErr) {
	if _, err := pd.orderRepository.FindOrderByID(orderID); err != nil {
		return nil, err
	}

	return pd.paymentRepository.FindPaymentsByOrderID(orderID)
}
