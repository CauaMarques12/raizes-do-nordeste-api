package repository

import (
	"context"
	"time"

	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/database/mongodb"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/configuration/rest_err"
	"github.com/CauaMarques12/raizes-do-nordeste-api/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type OrderRepository interface {
	CreateOrder(model.OrderDomainInterface) *rest_err.RestErr
	FindOrderByID(string) (model.OrderDomainInterface, *rest_err.RestErr)
	FindOrders(channel, status string, page, limit int64) ([]model.OrderDomainInterface, *rest_err.RestErr)
	UpdateStatus(orderID, status string) (model.OrderDomainInterface, *rest_err.RestErr)
}

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		collection: mongodb.GetCollection("orders"),
	}
}

func (or *orderRepository) CreateOrder(orderDomain model.OrderDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := or.collection.InsertOne(ctx, newOrderEntity(orderDomain))
	if err != nil {
		return rest_err.NewInternalServerError("Error trying to create order")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		orderDomain.SetID(objectID.Hex())
	}

	return nil
}

func (or *orderRepository) FindOrderByID(orderID string) (model.OrderDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid order id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity orderEntity
	if err := or.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Order not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find order")
	}

	return entity.toDomain(), nil
}

func (or *orderRepository) FindOrders(channel, status string, page, limit int64) ([]model.OrderDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if channel != "" {
		filter["channel"] = channel
	}
	if status != "" {
		filter["status"] = status
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	if page > 0 && limit > 0 {
		findOptions.SetSkip((page - 1) * limit)
		findOptions.SetLimit(limit)
	}

	cursor, err := or.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to find orders")
	}
	defer cursor.Close(ctx)

	var entities []orderEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to decode orders")
	}

	orders := make([]model.OrderDomainInterface, 0, len(entities))
	for _, entity := range entities {
		orders = append(orders, entity.toDomain())
	}

	return orders, nil
}

func (or *orderRepository) UpdateStatus(orderID, status string) (model.OrderDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid order id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updatedEntity orderEntity
	err = or.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"status": status, "updatedAt": time.Now().UTC()}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Order not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to update order status")
	}

	return updatedEntity.toDomain(), nil
}
