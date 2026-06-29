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

type PaymentRepository interface {
	CreatePayment(model.PaymentDomainInterface) *rest_err.RestErr
	FindPaymentByID(string) (model.PaymentDomainInterface, *rest_err.RestErr)
	FindPaymentsByOrderID(string, int64, int64) ([]model.PaymentDomainInterface, *rest_err.RestErr)
}

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository() PaymentRepository {
	repository := &paymentRepository{
		collection: mongodb.GetCollection("payments"),
	}
	repository.createIndexes()
	return repository
}

func (pr *paymentRepository) CreatePayment(paymentDomain model.PaymentDomainInterface) *rest_err.RestErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := pr.collection.InsertOne(ctx, newPaymentEntity(paymentDomain))
	if err != nil {
		return rest_err.NewInternalServerError("Erro ao tentar criar pagamento")
	}

	if objectID, ok := result.InsertedID.(bson.ObjectID); ok {
		paymentDomain.SetID(objectID.Hex())
	}

	return nil
}

func (pr *paymentRepository) FindPaymentByID(paymentID string) (model.PaymentDomainInterface, *rest_err.RestErr) {
	objectID, err := bson.ObjectIDFromHex(paymentID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("Pagamento invalido")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entity paymentEntity
	if err := pr.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&entity); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Pagamento nao encontrado")
		}
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar pagamento")
	}

	return entity.toDomain(), nil
}

func (pr *paymentRepository) FindPaymentsByOrderID(orderID string, page, limit int64) ([]model.PaymentDomainInterface, *rest_err.RestErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	if page > 0 && limit > 0 {
		findOptions.SetSkip((page - 1) * limit)
		findOptions.SetLimit(limit)
	}

	cursor, err := pr.collection.Find(
		ctx,
		bson.M{"orderId": orderID},
		findOptions,
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar buscar pagamentos")
	}
	defer cursor.Close(ctx)

	var entities []paymentEntity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, rest_err.NewInternalServerError("Erro ao tentar decodificar pagamentos")
	}

	payments := make([]model.PaymentDomainInterface, 0, len(entities))
	for _, entity := range entities {
		payments = append(payments, entity.toDomain())
	}

	return payments, nil
}

func (pr *paymentRepository) createIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, _ = pr.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "orderId", Value: 1}, {Key: "createdAt", Value: -1}},
	})
}
