package model

import "time"

type OrderItemDomain struct {
	ProductID      string
	Quantity       int64
	UnitPriceCents int64
	SubtotalCents  int64
}

type OrderDomainInterface interface {
	GetID() string
	GetClientID() string
	GetUnitID() string
	GetChannel() string
	GetPaymentMethod() string
	GetStatus() string
	GetTotalCents() int64
	GetItems() []OrderItemDomain
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetStatus(string)
	SetTotalCents(int64)
	SetItems([]OrderItemDomain)
}

func NewOrderDomain(
	clientID, unitID, channel, paymentMethod string,
	items []OrderItemDomain,
) OrderDomainInterface {
	now := time.Now().UTC()
	return &orderDomain{
		clientID:      clientID,
		unitID:        unitID,
		channel:       channel,
		paymentMethod: paymentMethod,
		status:        "AGUARDANDO_PAGAMENTO",
		items:         items,
		createdAt:     now,
		updatedAt:     now,
	}
}

func NewOrderDomainWithID(
	id, clientID, unitID, channel, paymentMethod, status string,
	totalCents int64,
	items []OrderItemDomain,
	createdAt, updatedAt time.Time,
) OrderDomainInterface {
	return &orderDomain{
		id:            id,
		clientID:      clientID,
		unitID:        unitID,
		channel:       channel,
		paymentMethod: paymentMethod,
		status:        status,
		totalCents:    totalCents,
		items:         items,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

type orderDomain struct {
	id            string
	clientID      string
	unitID        string
	channel       string
	paymentMethod string
	status        string
	totalCents    int64
	items         []OrderItemDomain
	createdAt     time.Time
	updatedAt     time.Time
}

func (od *orderDomain) GetID() string {
	return od.id
}

func (od *orderDomain) GetClientID() string {
	return od.clientID
}

func (od *orderDomain) GetUnitID() string {
	return od.unitID
}

func (od *orderDomain) GetChannel() string {
	return od.channel
}

func (od *orderDomain) GetPaymentMethod() string {
	return od.paymentMethod
}

func (od *orderDomain) GetStatus() string {
	return od.status
}

func (od *orderDomain) GetTotalCents() int64 {
	return od.totalCents
}

func (od *orderDomain) GetItems() []OrderItemDomain {
	return od.items
}

func (od *orderDomain) GetCreatedAt() time.Time {
	return od.createdAt
}

func (od *orderDomain) GetUpdatedAt() time.Time {
	return od.updatedAt
}

func (od *orderDomain) SetID(id string) {
	od.id = id
}

func (od *orderDomain) SetStatus(status string) {
	od.status = status
}

func (od *orderDomain) SetTotalCents(totalCents int64) {
	od.totalCents = totalCents
}

func (od *orderDomain) SetItems(items []OrderItemDomain) {
	od.items = items
}
