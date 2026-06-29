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

type LoyaltyRepository interface {
	CreateMovement(model.LoyaltyMovementDomainInterface) *rest_err.RestErr
	FindBalance(string) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr)
	FindMovements(string, int64, int64) ([]model.LoyaltyMovementDomainInterface, *rest_err.RestErr)
}

type loyaltyRepository struct {
	balanceCollection  *mongo.Collection
	movementCollection *mongo.Collection
}

func NewLoyaltyRepository() LoyaltyRepository {
	repository := &loyaltyRepository{
		balanceCollection:  mongodb.GetCollection("loyalty_balances"),
		movementCollection: mongodb.GetCollection("loyalty_movements"),
	}
	repository.createIndexes()
	return repository
}

func (lr *loyaltyRepository) CreateMovement(loyaltyMovementDomain model.LoyaltyMovementDomainInterface) *rest_err.RestErr {
	var balance model.LoyaltyBalanceDomainInterface
	var err *rest_err.RestErr

	switch loyaltyMovementDomain.GetType() {
	case "CREDITO":
		balance, err = lr.increasePoints(loyaltyMovementDomain.GetUserID(), loyaltyMovementDomain.GetPoints())
	case "RESGATE":
		balance, err = lr.decreasePoints(loyaltyMovementDomain.GetUserID(), loyaltyMovementDomain.GetPoints())
	default:
		return rest_err.NewBadRequestError("Tipo de movimento de fidelidade invalido")
	}
	if err != nil {
		return err
	}

	loyaltyMovementDomain.SetBalanceAfter(balance.GetPoints())
	movementEntity := newLoyaltyMovementEntity(loyaltyMovementDomain)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, mongoErr := lr.movementCollection.InsertOne(ctx, movementEntity)
	if mongoErr != nil {
		return rest_err.NewInternalServerError("Erro ao tentar criar movimento de fidelidade")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		loyaltyMovementDomain.SetID(objectID.Hex())
	}

	return nil
}

func (lr *loyaltyRepository) FindBalance(userID string) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity loyaltyBalanceEntity
	err := lr.balanceCollection.FindOne(ctx, bson.M{
		"userId": userID,
		"active": true,
	}).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Saldo de fidelidade nao encontrado")
		}
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar saldo de fidelidade")
	}

	return entity.toDomain(), nil
}

func (lr *loyaltyRepository) FindMovements(userID string, page, limit int64) ([]model.LoyaltyMovementDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	if page > 0 && limit > 0 {
		findOptions.SetSkip((page - 1) * limit)
		findOptions.SetLimit(limit)
	}

	cursor, err := lr.movementCollection.Find(
		ctx,
		bson.M{"userId": userID},
		findOptions,
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar historico de fidelidade")
	}
	defer cursor.Close(ctx)

	var entities []loyaltyMovementEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar decodificar historico de fidelidade")
	}

	movements := make([]model.LoyaltyMovementDomainInterface, 0, len(entities))
	for _, entity := range entities {
		movements = append(movements, entity.toDomain())
	}

	return movements, nil
}

func (lr *loyaltyRepository) increasePoints(userID string, points int64) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	now := time.Now().UTC()
	_, err := lr.balanceCollection.UpdateOne(
		ctx,
		bson.M{"userId": userID},
		bson.M{
			"$set": bson.M{
				"active":    true,
				"updatedAt": now,
			},
			"$setOnInsert": bson.M{
				"userId":    userID,
				"createdAt": now,
			},
			"$inc": bson.M{"points": points},
		},
		options.UpdateOne().SetUpsert(true),
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar adicionar pontos de fidelidade")
	}

	return lr.FindBalance(userID)
}

func (lr *loyaltyRepository) decreasePoints(userID string, points int64) (model.LoyaltyBalanceDomainInterface, *rest_err.RestErr) {
	currentBalance, err := lr.FindBalance(userID)
	if err != nil {
		if err.Code == 404 {
			return nil, rest_err.NewConflictError("Saldo de fidelidade insuficiente")
		}
		return nil, err
	}
	if currentBalance.GetPoints() < points {
		return nil, rest_err.NewConflictError("Saldo de fidelidade insuficiente")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, mongoErr := lr.balanceCollection.UpdateOne(
		ctx,
		bson.M{
			"userId": userID,
			"active": true,
			"points": bson.M{"$gte": points},
		},
		bson.M{
			"$set": bson.M{"updatedAt": time.Now().UTC()},
			"$inc": bson.M{"points": -points},
		},
	)
	if mongoErr != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar resgatar pontos de fidelidade")
	}
	if result.MatchedCount == 0 {
		return nil, rest_err.NewConflictError("Saldo de fidelidade insuficiente")
	}

	return lr.FindBalance(userID)
}

func (lr *loyaltyRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = lr.balanceCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "userId", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	_, _ = lr.movementCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "userId", Value: 1}, {Key: "createdAt", Value: -1}},
	})
}
