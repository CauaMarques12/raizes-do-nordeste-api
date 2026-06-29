package service

import (
	"fmt"
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
)

func (pd *paymentDomainService) ProcessPayment(paymentDomain model.PaymentDomainInterface) (model.PaymentDomainInterface, *rest_err.RestErr) {
	orderDomain, err := pd.orderRepository.FindOrderByID(paymentDomain.GetOrderID())
	if err != nil {
		return nil, err
	}

	if orderDomain.GetStatus() != "AGUARDANDO_PAGAMENTO" {
		return nil, rest_err.NewConflictError("Pedido nao permite pagamento")
	}

	if orderDomain.GetTotalCents() != paymentDomain.GetAmountCents() {
		return nil, rest_err.NewConflictError("Valor do pagamento diferente do total do pedido")
	}

	paymentDomain.SetGatewayTransactionID(fmt.Sprintf("mock-%d", time.Now().UTC().UnixNano()))
	if err := pd.paymentRepository.CreatePayment(paymentDomain); err != nil {
		return nil, err
	}

	if paymentDomain.GetStatus() == "APROVADO" {
		if _, err := pd.orderRepository.UpdateStatus(orderDomain.GetID(), "PAGO"); err != nil {
			return nil, err
		}
		if err := pd.addLoyaltyPoints(orderDomain); err != nil {
			return nil, err
		}
		return paymentDomain, nil
	}

	if err := pd.restoreOrderStock(orderDomain); err != nil {
		return nil, err
	}

	if _, err := pd.orderRepository.UpdateStatus(orderDomain.GetID(), "CANCELADO"); err != nil {
		return nil, err
	}

	return paymentDomain, nil
}

func (pd *paymentDomainService) addLoyaltyPoints(orderDomain model.OrderDomainInterface) *rest_err.RestErr {
	userDomain, err := pd.userRepository.FindUserByID(orderDomain.GetClientID())
	if err != nil {
		return err
	}

	if !userDomain.GetFidelidadeConsentida() {
		return nil
	}

	points := orderDomain.GetTotalCents() / 1000
	if points < 1 {
		points = 1
	}

	movement := model.NewLoyaltyMovementDomain(
		orderDomain.GetClientID(),
		"CREDITO",
		points,
		fmt.Sprintf("Pagamento aprovado do pedido %s", orderDomain.GetID()),
		orderDomain.GetID(),
	)

	return pd.loyaltyRepository.CreateMovement(movement)
}

func (pd *paymentDomainService) restoreOrderStock(orderDomain model.OrderDomainInterface) *rest_err.RestErr {
	for _, item := range orderDomain.GetItems() {
		movement := model.NewStockMovementDomain(
			orderDomain.GetUnitID(),
			item.ProductID,
			"ENTRADA",
			item.Quantity,
			fmt.Sprintf("Pagamento recusado do pedido %s", orderDomain.GetID()),
		)
		if err := pd.stockRepository.CreateMovement(movement); err != nil {
			return err
		}
	}

	return nil
}
