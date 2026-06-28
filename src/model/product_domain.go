package model

import "time"

type ProductDomainInterface interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetCategory() string
	GetPriceCents() int64
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
}

func NewProductDomain(name, description, category string, priceCents int64) ProductDomainInterface {
	now := time.Now().UTC()
	return &productDomain{
		name:        name,
		description: description,
		category:    category,
		priceCents:  priceCents,
		active:      true,
		createdAt:   now,
		updatedAt:   now,
	}
}

func NewProductUpdateDomain(name, description, category string, priceCents int64) ProductDomainInterface {
	return &productDomain{
		name:        name,
		description: description,
		category:    category,
		priceCents:  priceCents,
	}
}

func NewProductDomainWithID(
	id, name, description, category string,
	priceCents int64,
	active bool,
	createdAt, updatedAt time.Time,
) ProductDomainInterface {
	return &productDomain{
		id:          id,
		name:        name,
		description: description,
		category:    category,
		priceCents:  priceCents,
		active:      active,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

type productDomain struct {
	id          string
	name        string
	description string
	category    string
	priceCents  int64
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (pd *productDomain) GetID() string {
	return pd.id
}

func (pd *productDomain) GetName() string {
	return pd.name
}

func (pd *productDomain) GetDescription() string {
	return pd.description
}

func (pd *productDomain) GetCategory() string {
	return pd.category
}

func (pd *productDomain) GetPriceCents() int64 {
	return pd.priceCents
}

func (pd *productDomain) GetActive() bool {
	return pd.active
}

func (pd *productDomain) GetCreatedAt() time.Time {
	return pd.createdAt
}

func (pd *productDomain) GetUpdatedAt() time.Time {
	return pd.updatedAt
}

func (pd *productDomain) SetID(id string) {
	pd.id = id
}
