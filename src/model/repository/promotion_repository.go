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

type PromotionRepository interface {
	CreatePromotion(model.PromotionDomainInterface) *rest_err.RestErr
	FindPromotionByID(string) (model.PromotionDomainInterface, *rest_err.RestErr)
	FindPromotionByCode(string) (model.PromotionDomainInterface, *rest_err.RestErr)
	FindPromotions(*bool, int64, int64) ([]model.PromotionDomainInterface, *rest_err.RestErr)
	UpdatePromotion(string, model.PromotionDomainInterface) (model.PromotionDomainInterface, *rest_err.RestErr)
}

type promotionRepository struct {
	collection *mongo.Collection
}

func NewPromotionRepository() PromotionRepository {
	repository := &promotionRepository{
		collection: mongodb.GetCollection("promotions"),
	}
	repository.createIndexes()
	return repository
}

func (pr *promotionRepository) CreatePromotion(promotionDomain model.PromotionDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := pr.collection.InsertOne(ctx, newPromotionEntity(promotionDomain))
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return rest_err.NewConflictError("Promocao ja existe")
		}
		return rest_err.NewInternalServerError("Erro ao tentar criar promocao")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		promotionDomain.SetID(objectID.Hex())
	}

	return nil
}

func (pr *promotionRepository) FindPromotionByID(promotionID string) (model.PromotionDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(promotionID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Promocao invalida")
	}

	return pr.findOne(bson.M{"_id": objectID})
}

func (pr *promotionRepository) FindPromotionByCode(code string) (model.PromotionDomainInterface, *rest_err.RestErr) {
	return pr.findOne(bson.M{"code": code, "active": true})
}

func (pr *promotionRepository) FindPromotions(active *bool, page, limit int64) ([]model.PromotionDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if active != nil {
		filter["active"] = *active
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	if page > 0 && limit > 0 {
		findOptions.SetSkip((page - 1) * limit)
		findOptions.SetLimit(limit)
	}

	cursor, err := pr.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar promocoes")
	}
	defer cursor.Close(ctx)

	var entities []promotionEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar decodificar promocoes")
	}

	promotions := make([]model.PromotionDomainInterface, 0, len(entities))
	for _, entity := range entities {
		promotions = append(promotions, entity.toDomain())
	}

	return promotions, nil
}

func (pr *promotionRepository) UpdatePromotion(promotionID string, promotionDomain model.PromotionDomainInterface) (model.PromotionDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(promotionID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Promocao invalida")
	}

	updateFields := bson.M{}
	if promotionDomain.GetName() != "" {
		updateFields["name"] = promotionDomain.GetName()
	}
	if promotionDomain.GetCode() != "" {
		updateFields["code"] = promotionDomain.GetCode()
	}
	if promotionDomain.GetDiscountPercent() > 0 {
		updateFields["discountPercent"] = promotionDomain.GetDiscountPercent()
	}
	if promotionDomain.HasActive() {
		updateFields["active"] = promotionDomain.GetActive()
	}
	if len(updateFields) == 0 {
		return nil, rest_err.NewBadRequestError("No fields to update")
	}

	updateFields["updatedAt"] = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updatedEntity promotionEntity
	err = pr.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Promocao nao encontrada")
		}
		if mongo.IsDuplicateKeyError(err) {
			return nil, rest_err.NewConflictError("Promocao ja existe")
		}
		return nil, rest_err.NewInternalServerError("Erro ao tentar atualizar promocao")
	}

	return updatedEntity.toDomain(), nil
}

func (pr *promotionRepository) findOne(filter bson.M) (model.PromotionDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity promotionEntity
	if err := pr.collection.FindOne(ctx, filter).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Promocao nao encontrada")
		}
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar promocao")
	}

	return entity.toDomain(), nil
}

func (pr *promotionRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = pr.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "code", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
