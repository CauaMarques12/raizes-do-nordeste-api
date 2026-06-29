package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type loyaltyBalanceEntity struct {
	ID        *bson.ObjectID `bson:"_id,omitempty"`
	UserID    string         `bson:"userId"`
	Points    int64          `bson:"points"`
	Active    bool           `bson:"active"`
	CreatedAt time.Time      `bson:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt"`
}

type loyaltyMovementEntity struct {
	ID           *bson.ObjectID `bson:"_id,omitempty"`
	UserID       string         `bson:"userId"`
	Type         string         `bson:"type"`
	Points       int64          `bson:"points"`
	Reason       string         `bson:"reason"`
	OrderID      string         `bson:"orderId,omitempty"`
	BalanceAfter int64          `bson:"balanceAfter"`
	CreatedAt    time.Time      `bson:"createdAt"`
}

func newLoyaltyMovementEntity(loyaltyMovementDomain model.LoyaltyMovementDomainInterface) loyaltyMovementEntity {
	var id *bson.ObjectID
	if loyaltyMovementDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(loyaltyMovementDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return loyaltyMovementEntity{
		ID:           id,
		UserID:       loyaltyMovementDomain.GetUserID(),
		Type:         loyaltyMovementDomain.GetType(),
		Points:       loyaltyMovementDomain.GetPoints(),
		Reason:       loyaltyMovementDomain.GetReason(),
		OrderID:      loyaltyMovementDomain.GetOrderID(),
		BalanceAfter: loyaltyMovementDomain.GetBalanceAfter(),
		CreatedAt:    loyaltyMovementDomain.GetCreatedAt(),
	}
}

func (lbe loyaltyBalanceEntity) toDomain() model.LoyaltyBalanceDomainInterface {
	id := ""
	if lbe.ID != nil {
		id = lbe.ID.Hex()
	}

	return model.NewLoyaltyBalanceDomainWithID(
		id,
		lbe.UserID,
		lbe.Points,
		lbe.Active,
		lbe.CreatedAt,
		lbe.UpdatedAt,
	)
}

func (lme loyaltyMovementEntity) toDomain() model.LoyaltyMovementDomainInterface {
	id := ""
	if lme.ID != nil {
		id = lme.ID.Hex()
	}

	return model.NewLoyaltyMovementDomainWithID(
		id,
		lme.UserID,
		lme.Type,
		lme.Points,
		lme.Reason,
		lme.OrderID,
		lme.BalanceAfter,
		lme.CreatedAt,
	)
}
