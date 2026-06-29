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

type StockRepository interface {
	CreateMovement(model.StockMovementDomainInterface) *rest_err.RestErr
	FindBalance(unitID, productID string) (model.StockBalanceDomainInterface, *rest_err.RestErr)
}

type stockRepository struct {
	balanceCollection  *mongo.Collection
	movementCollection *mongo.Collection
}

func NewStockRepository() StockRepository {
	repository := &stockRepository{
		balanceCollection:  mongodb.GetCollection("stock_balances"),
		movementCollection: mongodb.GetCollection("stock_movements"),
	}
	repository.createIndexes()
	return repository
}

func (sr *stockRepository) CreateMovement(stockMovementDomain model.StockMovementDomainInterface) *rest_err.RestErr {
	var balance model.StockBalanceDomainInterface
	var err *rest_err.RestErr

	switch stockMovementDomain.GetType() {
	case "ENTRADA":
		balance, err = sr.increaseStock(stockMovementDomain.GetUnitID(), stockMovementDomain.GetProductID(), stockMovementDomain.GetQuantity())
	case "SAIDA":
		balance, err = sr.decreaseStock(stockMovementDomain.GetUnitID(), stockMovementDomain.GetProductID(), stockMovementDomain.GetQuantity())
	default:
		return rest_err.NewBadRequestError("Invalid stock movement type")
	}
	if err != nil {
		return err
	}

	stockMovementDomain.SetBalanceAfter(balance.GetQuantity())
	movementEntity := newStockMovementEntity(stockMovementDomain)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, mongoErr := sr.movementCollection.InsertOne(ctx, movementEntity)
	if mongoErr != nil {
		return rest_err.NewInternalServerError("Error trying to create stock movement")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		stockMovementDomain.SetID(objectID.Hex())
	}

	return nil
}

func (sr *stockRepository) FindBalance(unitID, productID string) (model.StockBalanceDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity stockBalanceEntity
	err := sr.balanceCollection.FindOne(ctx, bson.M{
		"unitId":    unitID,
		"productId": productID,
		"active":    true,
	}).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Stock balance not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find stock balance")
	}

	return entity.toDomain(), nil
}

func (sr *stockRepository) increaseStock(unitID, productID string, quantity int64) (model.StockBalanceDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now().UTC()
	_, err := sr.balanceCollection.UpdateOne(
		ctx,
		bson.M{"unitId": unitID, "productId": productID},
		bson.M{
			"$set": bson.M{
				"active":    true,
				"updatedAt": now,
			},
			"$setOnInsert": bson.M{
				"unitId":    unitID,
				"productId": productID,
				"createdAt": now,
			},
			"$inc": bson.M{"quantity": quantity},
		},
		options.UpdateOne().SetUpsert(true),
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to increase stock")
	}

	return sr.FindBalance(unitID, productID)
}

func (sr *stockRepository) decreaseStock(unitID, productID string, quantity int64) (model.StockBalanceDomainInterface, *rest_err.RestErr) {
	currentBalance, err := sr.FindBalance(unitID, productID)
	if err != nil {
		if err.Code == 404 {
			return nil, rest_err.NewConflictError("Stock is not available")
		}
		return nil, err
	}
	if currentBalance.GetQuantity() < quantity {
		return nil, rest_err.NewConflictError("Insufficient stock")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, mongoErr := sr.balanceCollection.UpdateOne(
		ctx,
		bson.M{
			"unitId":    unitID,
			"productId": productID,
			"active":    true,
			"quantity":  bson.M{"$gte": quantity},
		},
		bson.M{
			"$set": bson.M{"updatedAt": time.Now().UTC()},
			"$inc": bson.M{"quantity": -quantity},
		},
	)
	if mongoErr != nil {
		return nil, rest_err.NewInternalServerError("Error trying to decrease stock")
	}
	if result.MatchedCount == 0 {
		return nil, rest_err.NewConflictError("Insufficient stock")
	}

	return sr.FindBalance(unitID, productID)
}

func (sr *stockRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = sr.balanceCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "unitId", Value: 1}, {Key: "productId", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
