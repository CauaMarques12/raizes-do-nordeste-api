package model

import "time"

type LoyaltyBalanceDomainInterface interface {
	GetID() string
	GetUserID() string
	GetPoints() int64
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetPoints(int64)
}

type LoyaltyMovementDomainInterface interface {
	GetID() string
	GetUserID() string
	GetType() string
	GetPoints() int64
	GetReason() string
	GetOrderID() string
	GetBalanceAfter() int64
	GetCreatedAt() time.Time
	SetID(string)
	SetBalanceAfter(int64)
}

func NewLoyaltyBalanceDomain(
	userID string,
	points int64,
) LoyaltyBalanceDomainInterface {
	now := time.Now().UTC()
	return &loyaltyBalanceDomain{
		userID:    userID,
		points:    points,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}
}

func NewLoyaltyBalanceDomainWithID(
	id, userID string,
	points int64,
	active bool,
	createdAt, updatedAt time.Time,
) LoyaltyBalanceDomainInterface {
	return &loyaltyBalanceDomain{
		id:        id,
		userID:    userID,
		points:    points,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewLoyaltyMovementDomain(
	userID, movementType string,
	points int64,
	reason, orderID string,
) LoyaltyMovementDomainInterface {
	return &loyaltyMovementDomain{
		userID:       userID,
		movementType: movementType,
		points:       points,
		reason:       reason,
		orderID:      orderID,
		createdAt:    time.Now().UTC(),
	}
}

func NewLoyaltyMovementDomainWithID(
	id, userID, movementType string,
	points int64,
	reason, orderID string,
	balanceAfter int64,
	createdAt time.Time,
) LoyaltyMovementDomainInterface {
	return &loyaltyMovementDomain{
		id:           id,
		userID:       userID,
		movementType: movementType,
		points:       points,
		reason:       reason,
		orderID:      orderID,
		balanceAfter: balanceAfter,
		createdAt:    createdAt,
	}
}

type loyaltyBalanceDomain struct {
	id        string
	userID    string
	points    int64
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

type loyaltyMovementDomain struct {
	id           string
	userID       string
	movementType string
	points       int64
	reason       string
	orderID      string
	balanceAfter int64
	createdAt    time.Time
}

func (ld *loyaltyBalanceDomain) GetID() string {
	return ld.id
}

func (ld *loyaltyBalanceDomain) GetUserID() string {
	return ld.userID
}

func (ld *loyaltyBalanceDomain) GetPoints() int64 {
	return ld.points
}

func (ld *loyaltyBalanceDomain) GetActive() bool {
	return ld.active
}

func (ld *loyaltyBalanceDomain) GetCreatedAt() time.Time {
	return ld.createdAt
}

func (ld *loyaltyBalanceDomain) GetUpdatedAt() time.Time {
	return ld.updatedAt
}

func (ld *loyaltyBalanceDomain) SetID(id string) {
	ld.id = id
}

func (ld *loyaltyBalanceDomain) SetPoints(points int64) {
	ld.points = points
}

func (lm *loyaltyMovementDomain) GetID() string {
	return lm.id
}

func (lm *loyaltyMovementDomain) GetUserID() string {
	return lm.userID
}

func (lm *loyaltyMovementDomain) GetType() string {
	return lm.movementType
}

func (lm *loyaltyMovementDomain) GetPoints() int64 {
	return lm.points
}

func (lm *loyaltyMovementDomain) GetReason() string {
	return lm.reason
}

func (lm *loyaltyMovementDomain) GetOrderID() string {
	return lm.orderID
}

func (lm *loyaltyMovementDomain) GetBalanceAfter() int64 {
	return lm.balanceAfter
}

func (lm *loyaltyMovementDomain) GetCreatedAt() time.Time {
	return lm.createdAt
}

func (lm *loyaltyMovementDomain) SetID(id string) {
	lm.id = id
}

func (lm *loyaltyMovementDomain) SetBalanceAfter(balanceAfter int64) {
	lm.balanceAfter = balanceAfter
}
