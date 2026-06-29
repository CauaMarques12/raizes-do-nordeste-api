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

type ProductRepository interface {
	CreateProduct(model.ProductDomainInterface) *rest_err.RestErr
	FindProductByID(string) (model.ProductDomainInterface, *rest_err.RestErr)
	FindProducts(string, int64, int64) ([]model.ProductDomainInterface, *rest_err.RestErr)
	UpdateProduct(string, model.ProductDomainInterface) (model.ProductDomainInterface, *rest_err.RestErr)
}

type productRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		collection: mongodb.GetCollection("products"),
	}
}

func (pr *productRepository) CreateProduct(productDomain model.ProductDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := pr.collection.InsertOne(ctx, newProductEntity(productDomain))
	if err != nil {
		return rest_err.NewInternalServerError("Error trying to create product")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		productDomain.SetID(objectID.Hex())
	}

	return nil
}

func (pr *productRepository) FindProductByID(productID string) (model.ProductDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(productID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid product id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity productEntity
	if err := pr.collection.FindOne(ctx, bson.M{"_id": objectID, "active": true}).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Product not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find product")
	}

	return entity.toDomain(), nil
}

func (pr *productRepository) FindProducts(category string, page, limit int64) ([]model.ProductDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"active": true}
	if category != "" {
		filter["category"] = category
	}

	findOptions := options.Find()
	if page > 0 && limit > 0 {
		findOptions.SetSkip((page - 1) * limit)
		findOptions.SetLimit(limit)
	}

	cursor, err := pr.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to find products")
	}
	defer cursor.Close(ctx)

	var entities []productEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to decode products")
	}

	products := make([]model.ProductDomainInterface, 0, len(entities))
	for _, entity := range entities {
		products = append(products, entity.toDomain())
	}

	return products, nil
}

func (pr *productRepository) UpdateProduct(productID string, productDomain model.ProductDomainInterface) (model.ProductDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(productID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Invalid product id")
	}

	updateFields := bson.M{}
	if productDomain.GetName() != "" {
		updateFields["name"] = productDomain.GetName()
	}
	if productDomain.GetDescription() != "" {
		updateFields["description"] = productDomain.GetDescription()
	}
	if productDomain.GetCategory() != "" {
		updateFields["category"] = productDomain.GetCategory()
	}
	if productDomain.GetPriceCents() > 0 {
		updateFields["priceCents"] = productDomain.GetPriceCents()
	}
	if len(updateFields) == 0 {
		return nil, rest_err.NewBadRequestError("No fields to update")
	}

	updateFields["updatedAt"] = time.Now().UTC()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updatedEntity productEntity
	err = pr.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID, "active": true},
		bson.M{"$set": updateFields},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Product not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to update product")
	}

	return updatedEntity.toDomain(), nil
}
