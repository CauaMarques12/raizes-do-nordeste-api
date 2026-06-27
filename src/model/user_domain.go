package model

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetRole() string
	GetFidelidadeConsentida() bool
	HasFidelidadeConsentida() bool
	GetActive() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	SetID(string)
	SetPassword(string)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetActive(bool)
	EncryptPassword() *rest_err.RestErr
}

func NewUserDomain(
	email, password, name, role string,
	fidelidadeConsentida *bool,
) UserDomainInterface {
	now := time.Now().UTC()
	if role == "" {
		role = "CLIENTE"
	}
	return &userDomain{
		email:                email,
		password:             password,
		name:                 name,
		role:                 role,
		fidelidadeConsentida: fidelidadeConsentida,
		active:               true,
		createdAt:            now,
		updatedAt:            now,
	}
}

func NewUserUpdateDomain(
	email, password, name, role string,
	fidelidadeConsentida *bool,
) UserDomainInterface {
	return &userDomain{
		email:                email,
		password:             password,
		name:                 name,
		role:                 role,
		fidelidadeConsentida: fidelidadeConsentida,
	}
}

func NewUserDomainWithID(
	id, email, password, name, role string,
	fidelidadeConsentida bool,
	active bool,
	createdAt, updatedAt time.Time,
) UserDomainInterface {
	return &userDomain{
		id:                   id,
		email:                email,
		password:             password,
		name:                 name,
		role:                 role,
		fidelidadeConsentida: &fidelidadeConsentida,
		active:               active,
		createdAt:            createdAt,
		updatedAt:            updatedAt,
	}
}

type userDomain struct {
	id                   string
	email                string
	password             string
	name                 string
	role                 string
	fidelidadeConsentida *bool
	active               bool
	createdAt            time.Time
	updatedAt            time.Time
}

func (ud *userDomain) GetID() string {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetRole() string {
	return ud.role
}

func (ud *userDomain) GetFidelidadeConsentida() bool {
	if ud.fidelidadeConsentida == nil {
		return false
	}
	return *ud.fidelidadeConsentida
}

func (ud *userDomain) HasFidelidadeConsentida() bool {
	return ud.fidelidadeConsentida != nil
}

func (ud *userDomain) GetActive() bool {
	return ud.active
}

func (ud *userDomain) GetCreatedAt() time.Time {
	return ud.createdAt
}

func (ud *userDomain) GetUpdatedAt() time.Time {
	return ud.updatedAt
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

func (ud *userDomain) SetPassword(password string) {
	ud.password = password
}

func (ud *userDomain) SetCreatedAt(createdAt time.Time) {
	ud.createdAt = createdAt
}

func (ud *userDomain) SetUpdatedAt(updatedAt time.Time) {
	ud.updatedAt = updatedAt
}

func (ud *userDomain) SetActive(active bool) {
	ud.active = active
}

func (ud *userDomain) EncryptPassword() *rest_err.RestErr {
	hash, err := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	if err != nil {
		return rest_err.NewInternalServerError("Error trying to encrypt password")
	}

	ud.password = string(hash)
	return nil
}
