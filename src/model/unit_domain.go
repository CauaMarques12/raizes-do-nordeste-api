package model

import "time"

type UnitDomainInterface interface {
	GetID() string
	GetName() string
	GetAddress() string
	GetCity() string
	GetState() string
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
}

func NewUnitDomain(name, address, city, state string) UnitDomainInterface {
	now := time.Now().UTC()
	return &unitDomain{
		name:      name,
		address:   address,
		city:      city,
		state:     state,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}
}

func NewUnitUpdateDomain(name, address, city, state string) UnitDomainInterface {
	return &unitDomain{
		name:    name,
		address: address,
		city:    city,
		state:   state,
	}
}

func NewUnitDomainWithID(
	id, name, address, city, state string,
	active bool,
	createdAt, updatedAt time.Time,
) UnitDomainInterface {
	return &unitDomain{
		id:        id,
		name:      name,
		address:   address,
		city:      city,
		state:     state,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

type unitDomain struct {
	id        string
	name      string
	address   string
	city      string
	state     string
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func (ud *unitDomain) GetID() string {
	return ud.id
}

func (ud *unitDomain) GetName() string {
	return ud.name
}

func (ud *unitDomain) GetAddress() string {
	return ud.address
}

func (ud *unitDomain) GetCity() string {
	return ud.city
}

func (ud *unitDomain) GetState() string {
	return ud.state
}

func (ud *unitDomain) GetActive() bool {
	return ud.active
}

func (ud *unitDomain) GetCreatedAt() time.Time {
	return ud.createdAt
}

func (ud *unitDomain) GetUpdatedAt() time.Time {
	return ud.updatedAt
}

func (ud *unitDomain) SetID(id string) {
	ud.id = id
}
