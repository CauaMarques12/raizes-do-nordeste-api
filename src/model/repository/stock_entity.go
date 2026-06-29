package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type stockBalanceEntity struct {
	ID        *bson.ObjectID `bson:"_id,omitempty"`
	UnitID    string         `bson:"unitId"`
	ProductID string         `bson:"productId"`
	Quantity  int64          `bson:"quantity"`
	Active    bool           `bson:"active"`
	CreatedAt time.Time      `bson:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt"`
}

type stockMovementEntity struct {
	ID           *bson.ObjectID `bson:"_id,omitempty"`
	UnitID       string         `bson:"unitId"`
	ProductID    string         `bson:"productId"`
	Type         string         `bson:"type"`
	Quantity     int64          `bson:"quantity"`
	Reason       string         `bson:"reason"`
	BalanceAfter int64          `bson:"balanceAfter"`
	CreatedAt    time.Time      `bson:"createdAt"`
}

func (sbe stockBalanceEntity) toDomain() model.StockBalanceDomainInterface {
	id := ""
	if sbe.ID != nil {
		id = sbe.ID.Hex()
	}

	return model.NewStockBalanceDomainWithID(
		id,
		sbe.UnitID,
		sbe.ProductID,
		sbe.Quantity,
		sbe.Active,
		sbe.CreatedAt,
		sbe.UpdatedAt,
	)
}

func newStockMovementEntity(stockMovementDomain model.StockMovementDomainInterface) stockMovementEntity {
	var id *bson.ObjectID
	if stockMovementDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(stockMovementDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return stockMovementEntity{
		ID:           id,
		UnitID:       stockMovementDomain.GetUnitID(),
		ProductID:    stockMovementDomain.GetProductID(),
		Type:         stockMovementDomain.GetType(),
		Quantity:     stockMovementDomain.GetQuantity(),
		Reason:       stockMovementDomain.GetReason(),
		BalanceAfter: stockMovementDomain.GetBalanceAfter(),
		CreatedAt:    stockMovementDomain.GetCreatedAt(),
	}
}

func (sme stockMovementEntity) toDomain() model.StockMovementDomainInterface {
	id := ""
	if sme.ID != nil {
		id = sme.ID.Hex()
	}

	return model.NewStockMovementDomainWithID(
		id,
		sme.UnitID,
		sme.ProductID,
		sme.Type,
		sme.Quantity,
		sme.Reason,
		sme.BalanceAfter,
		sme.CreatedAt,
	)
}
