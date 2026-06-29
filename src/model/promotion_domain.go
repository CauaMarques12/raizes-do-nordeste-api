package model

import "time"

type PromotionDomainInterface interface {
	GetID() string
	GetName() string
	GetCode() string
	GetDiscountPercent() int64
	GetActive() bool
	HasActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetCode(string)
}

func NewPromotionDomain(
	name, code string,
	discountPercent int64,
	active *bool,
) PromotionDomainInterface {
	now := time.Now().UTC()
	if active == nil {
		defaultActive := true
		active = &defaultActive
	}
	return &promotionDomain{
		name:            name,
		code:            code,
		discountPercent: discountPercent,
		active:          active,
		createdAt:       now,
		updatedAt:       now,
	}
}

func NewPromotionUpdateDomain(
	name, code string,
	discountPercent int64,
	active *bool,
) PromotionDomainInterface {
	return &promotionDomain{
		name:            name,
		code:            code,
		discountPercent: discountPercent,
		active:          active,
	}
}

func NewPromotionDomainWithID(
	id, name, code string,
	discountPercent int64,
	active bool,
	createdAt, updatedAt time.Time,
) PromotionDomainInterface {
	return &promotionDomain{
		id:              id,
		name:            name,
		code:            code,
		discountPercent: discountPercent,
		active:          &active,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
}

type promotionDomain struct {
	id              string
	name            string
	code            string
	discountPercent int64
	active          *bool
	createdAt       time.Time
	updatedAt       time.Time
}

func (pd *promotionDomain) GetID() string {
	return pd.id
}

func (pd *promotionDomain) GetName() string {
	return pd.name
}

func (pd *promotionDomain) GetCode() string {
	return pd.code
}

func (pd *promotionDomain) GetDiscountPercent() int64 {
	return pd.discountPercent
}

func (pd *promotionDomain) GetActive() bool {
	if pd.active == nil {
		return false
	}
	return *pd.active
}

func (pd *promotionDomain) HasActive() bool {
	return pd.active != nil
}

func (pd *promotionDomain) GetCreatedAt() time.Time {
	return pd.createdAt
}

func (pd *promotionDomain) GetUpdatedAt() time.Time {
	return pd.updatedAt
}

func (pd *promotionDomain) SetID(id string) {
	pd.id = id
}

func (pd *promotionDomain) SetCode(code string) {
	pd.code = code
}
