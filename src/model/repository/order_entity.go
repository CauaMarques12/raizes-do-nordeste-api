package repository

import (
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type orderItemEntity struct {
	ProductID      string `bson:"productId"`
	Quantity       int64  `bson:"quantity"`
	UnitPriceCents int64  `bson:"unitPriceCents"`
	SubtotalCents  int64  `bson:"subtotalCents"`
}

type orderEntity struct {
	ID            *bson.ObjectID    `bson:"_id,omitempty"`
	ClientID      string            `bson:"clientId"`
	UnitID        string            `bson:"unitId"`
	Channel       string            `bson:"channel"`
	PaymentMethod string            `bson:"paymentMethod"`
	PromotionCode string            `bson:"promotionCode,omitempty"`
	Status        string            `bson:"status"`
	TotalCents    int64             `bson:"totalCents"`
	DiscountCents int64             `bson:"discountCents"`
	Items         []orderItemEntity `bson:"items"`
	CreatedAt     time.Time         `bson:"createdAt"`
	UpdatedAt     time.Time         `bson:"updatedAt"`
}

func newOrderEntity(orderDomain model.OrderDomainInterface) orderEntity {
	var id *bson.ObjectID
	if orderDomain.GetID() != "" {
		objectID, err := bson.ObjectIDFromHex(orderDomain.GetID())
		if err == nil {
			id = &objectID
		}
	}

	return orderEntity{
		ID:            id,
		ClientID:      orderDomain.GetClientID(),
		UnitID:        orderDomain.GetUnitID(),
		Channel:       orderDomain.GetChannel(),
		PaymentMethod: orderDomain.GetPaymentMethod(),
		PromotionCode: orderDomain.GetPromotionCode(),
		Status:        orderDomain.GetStatus(),
		TotalCents:    orderDomain.GetTotalCents(),
		DiscountCents: orderDomain.GetDiscountCents(),
		Items:         newOrderItemEntities(orderDomain.GetItems()),
		CreatedAt:     orderDomain.GetCreatedAt(),
		UpdatedAt:     orderDomain.GetUpdatedAt(),
	}
}

func (oe orderEntity) toDomain() model.OrderDomainInterface {
	id := ""
	if oe.ID != nil {
		id = oe.ID.Hex()
	}

	return model.NewOrderDomainWithID(
		id,
		oe.ClientID,
		oe.UnitID,
		oe.Channel,
		oe.PaymentMethod,
		oe.PromotionCode,
		oe.Status,
		oe.TotalCents,
		oe.DiscountCents,
		oe.toDomainItems(),
		oe.CreatedAt,
		oe.UpdatedAt,
	)
}

func newOrderItemEntities(items []model.OrderItemDomain) []orderItemEntity {
	entities := make([]orderItemEntity, 0, len(items))
	for _, item := range items {
		entities = append(entities, orderItemEntity{
			ProductID:      item.ProductID,
			Quantity:       item.Quantity,
			UnitPriceCents: item.UnitPriceCents,
			SubtotalCents:  item.SubtotalCents,
		})
	}

	return entities
}

func (oe orderEntity) toDomainItems() []model.OrderItemDomain {
	items := make([]model.OrderItemDomain, 0, len(oe.Items))
	for _, item := range oe.Items {
		items = append(items, model.OrderItemDomain{
			ProductID:      item.ProductID,
			Quantity:       item.Quantity,
			UnitPriceCents: item.UnitPriceCents,
			SubtotalCents:  item.SubtotalCents,
		})
	}

	return items
}
