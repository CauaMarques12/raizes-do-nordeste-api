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
	GetPromotionCode() string
	GetDiscountCents() int64
	GetStatus() string
	GetTotalCents() int64
	GetItems() []OrderItemDomain
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetStatus(string)
	SetTotalCents(int64)
	SetDiscountCents(int64)
	SetPromotionCode(string)
	SetItems([]OrderItemDomain)
}

func NewOrderDomain(
	clientID, unitID, channel, paymentMethod, promotionCode string,
	items []OrderItemDomain,
) OrderDomainInterface {
	now := time.Now().UTC()
	return &orderDomain{
		clientID:      clientID,
		unitID:        unitID,
		channel:       channel,
		paymentMethod: paymentMethod,
		promotionCode: promotionCode,
		status:        "AGUARDANDO_PAGAMENTO",
		items:         items,
		createdAt:     now,
		updatedAt:     now,
	}
}

func NewOrderDomainWithID(
	id, clientID, unitID, channel, paymentMethod, promotionCode, status string,
	totalCents, discountCents int64,
	items []OrderItemDomain,
	createdAt, updatedAt time.Time,
) OrderDomainInterface {
	return &orderDomain{
		id:            id,
		clientID:      clientID,
		unitID:        unitID,
		channel:       channel,
		paymentMethod: paymentMethod,
		promotionCode: promotionCode,
		status:        status,
		totalCents:    totalCents,
		discountCents: discountCents,
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
	promotionCode string
	status        string
	totalCents    int64
	discountCents int64
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

func (od *orderDomain) GetPromotionCode() string {
	return od.promotionCode
}

func (od *orderDomain) GetDiscountCents() int64 {
	return od.discountCents
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

func (od *orderDomain) SetDiscountCents(discountCents int64) {
	od.discountCents = discountCents
}

func (od *orderDomain) SetPromotionCode(promotionCode string) {
	od.promotionCode = promotionCode
}

func (od *orderDomain) SetItems(items []OrderItemDomain) {
	od.items = items
}
