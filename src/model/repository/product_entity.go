package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type productEntity struct {
	ID          *bson.ObjectID `bson:"_id,omitempty"`
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	Category    string         `bson:"category"`
	PriceCents  int64          `bson:"priceCents"`
	Active      bool           `bson:"active"`
	CreatedAt   time.Time      `bson:"createdAt"`
	UpdatedAt   time.Time      `bson:"updatedAt"`
}

func newProductEntity(productDomain model.ProductDomainInterface) productEntity {
	var id *bson.ObjectID
	if productDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(productDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return productEntity{
		ID:          id,
		Name:        productDomain.GetName(),
		Description: productDomain.GetDescription(),
		Category:    productDomain.GetCategory(),
		PriceCents:  productDomain.GetPriceCents(),
		Active:      productDomain.GetActive(),
		CreatedAt:   productDomain.GetCreatedAt(),
		UpdatedAt:   productDomain.GetUpdatedAt(),
	}
}

func (pe productEntity) toDomain() model.ProductDomainInterface {
	id := ""
	if pe.ID != nil {
		id = pe.ID.Hex()
	}

	return model.NewProductDomainWithID(
		id,
		pe.Name,
		pe.Description,
		pe.Category,
		pe.PriceCents,
		pe.Active,
		pe.CreatedAt,
		pe.UpdatedAt,
	)
}
