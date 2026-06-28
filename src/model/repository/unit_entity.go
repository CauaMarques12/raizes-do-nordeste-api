package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type unitEntity struct {
	ID        *bson.ObjectID `bson:"_id,omitempty"`
	Name      string         `bson:"name"`
	Address   string         `bson:"address"`
	City      string         `bson:"city"`
	State     string         `bson:"state"`
	Active    bool           `bson:"active"`
	CreatedAt time.Time      `bson:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt"`
}

func newUnitEntity(unitDomain model.UnitDomainInterface) unitEntity {
	var id *bson.ObjectID
	if unitDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(unitDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return unitEntity{
		ID:        id,
		Name:      unitDomain.GetName(),
		Address:   unitDomain.GetAddress(),
		City:      unitDomain.GetCity(),
		State:     unitDomain.GetState(),
		Active:    unitDomain.GetActive(),
		CreatedAt: unitDomain.GetCreatedAt(),
		UpdatedAt: unitDomain.GetUpdatedAt(),
	}
}

func (ue unitEntity) toDomain() model.UnitDomainInterface {
	id := ""
	if ue.ID != nil {
		id = ue.ID.Hex()
	}

	return model.NewUnitDomainWithID(
		id,
		ue.Name,
		ue.Address,
		ue.City,
		ue.State,
		ue.Active,
		ue.CreatedAt,
		ue.UpdatedAt,
	)
}
