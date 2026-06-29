package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type paymentEntity struct {
	ID                   *bson.ObjectID `bson:"_id,omitempty"`
	OrderID              string         `bson:"orderId"`
	Method               string         `bson:"method"`
	AmountCents          int64          `bson:"amountCents"`
	Status               string         `bson:"status"`
	GatewayTransactionID string         `bson:"gatewayTransactionId"`
	Message              string         `bson:"message"`
	CreatedAt            time.Time      `bson:"createdAt"`
	UpdatedAt            time.Time      `bson:"updatedAt"`
}

func newPaymentEntity(paymentDomain model.PaymentDomainInterface) paymentEntity {
	var id *bson.ObjectID
	if paymentDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(paymentDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return paymentEntity{
		ID:                   id,
		OrderID:              paymentDomain.GetOrderID(),
		Method:               paymentDomain.GetMethod(),
		AmountCents:          paymentDomain.GetAmountCents(),
		Status:               paymentDomain.GetStatus(),
		GatewayTransactionID: paymentDomain.GetGatewayTransactionID(),
		Message:              paymentDomain.GetMessage(),
		CreatedAt:            paymentDomain.GetCreatedAt(),
		UpdatedAt:            paymentDomain.GetUpdatedAt(),
	}
}

func (pe paymentEntity) toDomain() model.PaymentDomainInterface {
	id := ""
	if pe.ID != nil {
		id = pe.ID.Hex()
	}

	return model.NewPaymentDomainWithID(
		id,
		pe.OrderID,
		pe.Method,
		pe.Status,
		pe.GatewayTransactionID,
		pe.Message,
		pe.AmountCents,
		pe.CreatedAt,
		pe.UpdatedAt,
	)
}
