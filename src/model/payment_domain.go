package model

import "time"

type PaymentDomainInterface interface {
	GetID() string
	GetOrderID() string
	GetMethod() string
	GetAmountCents() int64
	GetStatus() string
	GetGatewayTransactionID() string
	GetMessage() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetStatus(string)
	SetGatewayTransactionID(string)
	SetMessage(string)
}

func NewPaymentDomain(
	orderID string,
	amountCents int64,
	approved bool,
) PaymentDomainInterface {
	now := time.Now().UTC()
	status := "RECUSADO"
	message := "Pagamento recusado pelo gateway mock"
	if approved {
		status = "APROVADO"
		message = "Pagamento aprovado pelo gateway mock"
	}

	return &paymentDomain{
		orderID:     orderID,
		method:      "MOCK",
		amountCents: amountCents,
		status:      status,
		message:     message,
		createdAt:   now,
		updatedAt:   now,
	}
}

func NewPaymentDomainWithID(
	id, orderID, method, status, gatewayTransactionID, message string,
	amountCents int64,
	createdAt, updatedAt time.Time,
) PaymentDomainInterface {
	return &paymentDomain{
		id:                   id,
		orderID:              orderID,
		method:               method,
		amountCents:          amountCents,
		status:               status,
		gatewayTransactionID: gatewayTransactionID,
		message:              message,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
	}
}

type paymentDomain struct {
	id                   string
	orderID              string
	method               string
	amountCents          int64
	status               string
	gatewayTransactionID string
	message              string
	createdAt            time.Time
	updatedAt            time.Time
}

func (pd *paymentDomain) GetID() string {
	return pd.id
}

func (pd *paymentDomain) GetOrderID() string {
	return pd.orderID
}

func (pd *paymentDomain) GetMethod() string {
	return pd.method
}

func (pd *paymentDomain) GetAmountCents() int64 {
	return pd.amountCents
}

func (pd *paymentDomain) GetStatus() string {
	return pd.status
}

func (pd *paymentDomain) GetGatewayTransactionID() string {
	return pd.gatewayTransactionID
}

func (pd *paymentDomain) GetMessage() string {
	return pd.message
}

func (pd *paymentDomain) GetCreatedAt() time.Time {
	return pd.createdAt
}

func (pd *paymentDomain) GetUpdatedAt() time.Time {
	return pd.updatedAt
}

func (pd *paymentDomain) SetID(id string) {
	pd.id = id
}

func (pd *paymentDomain) SetStatus(status string) {
	pd.status = status
}

func (pd *paymentDomain) SetGatewayTransactionID(gatewayTransactionID string) {
	pd.gatewayTransactionID = gatewayTransactionID
}

func (pd *paymentDomain) SetMessage(message string) {
	pd.message = message
}
