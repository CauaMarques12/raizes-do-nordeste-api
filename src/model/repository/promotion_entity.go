package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type promotionEntity struct {
	ID              *bson.ObjectID `bson:"_id,omitempty"`
	Name            string         `bson:"name"`
	Code            string         `bson:"code"`
	DiscountPercent int64          `bson:"discountPercent"`
	Active          bool           `bson:"active"`
	CreatedAt       time.Time      `bson:"createdAt"`
	UpdatedAt       time.Time      `bson:"updatedAt"`
}

func newPromotionEntity(promotionDomain model.PromotionDomainInterface) promotionEntity {
	var id *bson.ObjectID
	if promotionDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(promotionDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return promotionEntity{
		ID:              id,
		Name:            promotionDomain.GetName(),
		Code:            promotionDomain.GetCode(),
		DiscountPercent: promotionDomain.GetDiscountPercent(),
		Active:          promotionDomain.GetActive(),
		CreatedAt:       promotionDomain.GetCreatedAt(),
		UpdatedAt:       promotionDomain.GetUpdatedAt(),
	}
}

func (pe promotionEntity) toDomain() model.PromotionDomainInterface {
	id := ""
	if pe.ID != nil {
		id = pe.ID.Hex()
	}

	return model.NewPromotionDomainWithID(
		id,
		pe.Name,
		pe.Code,
		pe.DiscountPercent,
		pe.Active,
		pe.CreatedAt,
		pe.UpdatedAt,
	)
}
