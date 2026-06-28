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

type UnitRepository interface {
	CreateUnit(model.UnitDomainInterface) *rest_err.RestErr
	FindUnitByID(string) (model.UnitDomainInterface, *rest_err.RestErr)
	FindUnits() ([]model.UnitDomainInterface, *rest_err.RestErr)
	UpdateUnit(string, model.UnitDomainInterface) (model.UnitDomainInterface, *rest_err.RestErr)
}

type unitRepository struct {
	collection *mongo.Collection
}

func NewUnitRepository() UnitRepository {
	return &unitRepository{
		collection: mongodb.GetCollection("units"),
	}
}

func (ur *unitRepository) CreateUnit(unitDomain model.UnitDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := ur.collection.InsertOne(ctx, newUnitEntity(unitDomain))
	if err != nil {
		return rest_err.NewInternalServerError("Error trying to create unit")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		unitDomain.SetID(objectID.Hex())
	}

	return nil
}

func (ur *unitRepository) FindUnitByID(unitID string) (model.UnitDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(unitID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid unit id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity unitEntity
	if err := ur.collection.FindOne(ctx, bson.M{"_id": objectID, "active": true}).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Unit not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find unit")
	}

	return entity.toDomain(), nil
}

func (ur *unitRepository) FindUnits() ([]model.UnitDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := ur.collection.Find(ctx, bson.M{"active": true})
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to find units")
	}
	defer cursor.Close(ctx)

	var entities []unitEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to decode units")
	}

	units := make([]model.UnitDomainInterface, 0, len(entities))
	for _, entity := range entities {
		units = append(units, entity.toDomain())
	}

	return units, nil
}

func (ur *unitRepository) UpdateUnit(unitID string, unitDomain model.UnitDomainInterface) (model.UnitDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(unitID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid unit id")
	}

	updateFields := bson.M{}
	if unitDomain.GetName() != "" {
		updateFields["name"] = unitDomain.GetName()
	}
	if unitDomain.GetAddress() != "" {
		updateFields["address"] = unitDomain.GetAddress()
	}
	if unitDomain.GetCity() != "" {
		updateFields["city"] = unitDomain.GetCity()
	}
	if unitDomain.GetState() != "" {
		updateFields["state"] = unitDomain.GetState()
	}
	if len(updateFields) == 0 {
		return nil, rest_err.NewBadRequestError("No fields to update")
	}

	updateFields["updatedAt"] = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updatedEntity unitEntity
	err = ur.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID, "active": true},
		bson.M{"$set": updateFields},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Unit not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to update unit")
	}

	return updatedEntity.toDomain(), nil
}
