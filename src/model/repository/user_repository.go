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

type UserRepository interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() UserRepository {
	repository := &userRepository{
		collection: mongodb.GetCollection("users"),
	}
	repository.createIndexes()
	return repository
}

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	entity := newUserEntity(userDomain)
	result, err := ur.collection.InsertOne(ctx, entity)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return rest_err.NewConflictError("User already exists")
		}
		return rest_err.NewInternalServerError("Error trying to create user")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		userDomain.SetID(objectID.Hex())
	}

	return nil
}

func (ur *userRepository) FindUserByID(userID string) (model.UserDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid user id")
	}

	return ur.findOne(bson.M{"_id": objectID, "active": true})
}

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	return ur.findOne(bson.M{"email": email, "active": true})
}

func (ur *userRepository) UpdateUser(userID string, userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid user id")
	}

	updateFields := bson.M{}
	if userDomain.GetEmail() != "" {
		updateFields["email"] = userDomain.GetEmail()
	}
	if userDomain.GetPassword() != "" {
		updateFields["password"] = userDomain.GetPassword()
	}
	if userDomain.GetName() != "" {
		updateFields["name"] = userDomain.GetName()
	}
	if userDomain.GetRole() != "" {
		updateFields["role"] = userDomain.GetRole()
	}
	if userDomain.HasFidelidadeConsentida() {
		updateFields["fidelidadeConsentida"] = userDomain.GetFidelidadeConsentida()
	}
	if len(updateFields) == 0 {
		return nil, rest_err.NewBadRequestError("No fields to update")
	}

	updateFields["updatedAt"] = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updatedEntity userEntity
	err = ur.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID, "active": true},
		bson.M{"$set": updateFields},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		if mongo.IsDuplicateKeyError(err) {
			return nil, rest_err.NewConflictError("User already exists")
		}
		return nil, rest_err.NewInternalServerError("Error trying to update user")
	}

	return updatedEntity.toDomain(), nil
}

func (ur *userRepository) DeleteUser(userID string) *rest_err.RestErr {
	objectID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid user id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := ur.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID, "active": true},
		bson.M{"$set": bson.M{"active": false, "updatedAt": time.Now().UTC()}},
	)
	if err != nil {
		return rest_err.NewInternalServerError("Error trying to delete user")
	}
	if result.MatchedCount == 0 {
		return rest_err.NewNotFoundError("User not found")
	}

	return nil
}

func (ur *userRepository) findOne(filter bson.M) (model.UserDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity userEntity
	if err := ur.collection.FindOne(ctx, filter).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user")
	}

	return entity.toDomain(), nil
}

func (ur *userRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = ur.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
