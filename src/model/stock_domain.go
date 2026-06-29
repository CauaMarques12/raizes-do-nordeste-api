package model

import "time"

type StockBalanceDomainInterface interface {
	GetID() string
	GetUnitID() string
	GetProductID() string
	GetQuantity() int64
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetQuantity(int64)
}

type StockMovementDomainInterface interface {
	GetID() string
	GetUnitID() string
	GetProductID() string
	GetType() string
	GetQuantity() int64
	GetReason() string
	GetBalanceAfter() int64
	GetCreatedAt() time.Time
	SetID(string)
	SetBalanceAfter(int64)
}

func NewStockBalanceDomain(
	unitID, productID string,
	quantity int64,
) StockBalanceDomainInterface {
	now := time.Now().UTC()
	return &stockBalanceDomain{
		unitID:    unitID,
		productID: productID,
		quantity:  quantity,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}
}

func NewStockBalanceDomainWithID(
	id, unitID, productID string,
	quantity int64,
	active bool,
	createdAt, updatedAt time.Time,
) StockBalanceDomainInterface {
	return &stockBalanceDomain{
		id:        id,
		unitID:    unitID,
		productID: productID,
		quantity:  quantity,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewStockMovementDomain(
	unitID, productID, movementType string,
	quantity int64,
	reason string,
) StockMovementDomainInterface {
	return &stockMovementDomain{
		unitID:       unitID,
		productID:    productID,
		movementType: movementType,
		quantity:     quantity,
		reason:       reason,
		createdAt:    time.Now().UTC(),
	}
}

func NewStockMovementDomainWithID(
	id, unitID, productID, movementType string,
	quantity int64,
	reason string,
	balanceAfter int64,
	createdAt time.Time,
) StockMovementDomainInterface {
	return &stockMovementDomain{
		id:           id,
		unitID:       unitID,
		productID:    productID,
		movementType: movementType,
		quantity:     quantity,
		reason:       reason,
		balanceAfter: balanceAfter,
		createdAt:    createdAt,
	}
}

type stockBalanceDomain struct {
	id        string
	unitID    string
	productID string
	quantity  int64
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

type stockMovementDomain struct {
	id           string
	unitID       string
	productID    string
	movementType string
	quantity     int64
	reason       string
	balanceAfter int64
	createdAt    time.Time
}

func (sd *stockBalanceDomain) GetID() string {
	return sd.id
}

func (sd *stockBalanceDomain) GetUnitID() string {
	return sd.unitID
}

func (sd *stockBalanceDomain) GetProductID() string {
	return sd.productID
}

func (sd *stockBalanceDomain) GetQuantity() int64 {
	return sd.quantity
}

func (sd *stockBalanceDomain) GetActive() bool {
	return sd.active
}

func (sd *stockBalanceDomain) GetCreatedAt() time.Time {
	return sd.createdAt
}

func (sd *stockBalanceDomain) GetUpdatedAt() time.Time {
	return sd.updatedAt
}

func (sd *stockBalanceDomain) SetID(id string) {
	sd.id = id
}

func (sd *stockBalanceDomain) SetQuantity(quantity int64) {
	sd.quantity = quantity
}

func (sm *stockMovementDomain) GetID() string {
	return sm.id
}

func (sm *stockMovementDomain) GetUnitID() string {
	return sm.unitID
}

func (sm *stockMovementDomain) GetProductID() string {
	return sm.productID
}

func (sm *stockMovementDomain) GetType() string {
	return sm.movementType
}

func (sm *stockMovementDomain) GetQuantity() int64 {
	return sm.quantity
}

func (sm *stockMovementDomain) GetReason() string {
	return sm.reason
}

func (sm *stockMovementDomain) GetBalanceAfter() int64 {
	return sm.balanceAfter
}

func (sm *stockMovementDomain) GetCreatedAt() time.Time {
	return sm.createdAt
}

func (sm *stockMovementDomain) SetID(id string) {
	sm.id = id
}

func (sm *stockMovementDomain) SetBalanceAfter(balanceAfter int64) {
	sm.balanceAfter = balanceAfter
}
