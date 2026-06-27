package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type userEntity struct {
	ID                   *bson.ObjectID `bson:"_id,omitempty"`
	Email                string         `bson:"email"`
	Password             string         `bson:"password"`
	Name                 string         `bson:"name"`
	Role                 string         `bson:"role"`
	FidelidadeConsentida bool           `bson:"fidelidadeConsentida"`
	Active               bool           `bson:"active"`
	CreatedAt            time.Time      `bson:"createdAt"`
	UpdatedAt            time.Time      `bson:"updatedAt"`
}

func newUserEntity(userDomain model.UserDomainInterface) userEntity {
	var id *bson.ObjectID
	if userDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(userDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}
	return userEntity{
		ID:                   id,
		Email:                userDomain.GetEmail(),
		Password:             userDomain.GetPassword(),
		Name:                 userDomain.GetName(),
		Role:                 userDomain.GetRole(),
		FidelidadeConsentida: userDomain.GetFidelidadeConsentida(),
		Active:               userDomain.GetActive(),
		CreatedAt:            userDomain.GetCreatedAt(),
		UpdatedAt:            userDomain.GetUpdatedAt(),
	}
}

func (ue userEntity) toDomain() model.UserDomainInterface {
	id := ""
	if ue.ID != nil {
		id = ue.ID.Hex()
	}
	return model.NewUserDomainWithID(
		id,
		ue.Email,
		ue.Password,
		ue.Name,
		ue.Role,
		ue.FidelidadeConsentida,
		ue.Active,
		ue.CreatedAt,
		ue.UpdatedAt,
	)
}
