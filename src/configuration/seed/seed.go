package seed

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const (
	AdminID          = "660000000000000000000001"
	ClientID         = "660000000000000000000002"
	UnitID           = "660000000000000000000101"
	ProductCuscuzID  = "660000000000000000000201"
	ProductTapiocaID = "660000000000000000000202"
	ProductSucoID    = "660000000000000000000203"
	PromotionID      = "660000000000000000000301"
	AdminEmail       = "admin@raizes.dev"
	AdminPassword    = "Admin@123"
	ClientEmail      = "cliente@raizes.dev"
	ClientPassword   = "Cliente@123"
	PromotionCode    = "NORDESTE10"
)

func Run(ctx context.Context, database *mongo.Database) error {
	now := time.Now().UTC()

	if err := createIndexes(ctx, database); err != nil {
		return err
	}

	adminHash, err := hashPassword(AdminPassword)
	if err != nil {
		return err
	}
	clientHash, err := hashPassword(ClientPassword)
	if err != nil {
		return err
	}

	if err := upsertByID(ctx, database.Collection("users"), AdminID, bson.M{
		"email":                AdminEmail,
		"password":             adminHash,
		"name":                 "Administrador Seed",
		"role":                 "ADMIN",
		"fidelidadeConsentida": false,
		"active":               true,
		"updatedAt":            now,
	}, bson.M{"createdAt": now}); err != nil {
		return err
	}

	if err := upsertByID(ctx, database.Collection("users"), ClientID, bson.M{
		"email":                ClientEmail,
		"password":             clientHash,
		"name":                 "Cliente Seed",
		"role":                 "CLIENTE",
		"fidelidadeConsentida": true,
		"active":               true,
		"updatedAt":            now,
	}, bson.M{"createdAt": now}); err != nil {
		return err
	}

	if err := upsertByID(ctx, database.Collection("units"), UnitID, bson.M{
		"name":      "Unidade Recife Antigo",
		"address":   "Rua do Bom Jesus, 100",
		"city":      "Recife",
		"state":     "PE",
		"active":    true,
		"updatedAt": now,
	}, bson.M{"createdAt": now}); err != nil {
		return err
	}

	products := []struct {
		id          string
		name        string
		description string
		category    string
		priceCents  int64
	}{
		{ProductCuscuzID, "Cuscuz Nordestino", "Cuscuz com manteiga de garrafa", "comida", 1800},
		{ProductTapiocaID, "Tapioca de Queijo Coalho", "Tapioca recheada com queijo coalho", "comida", 2200},
		{ProductSucoID, "Suco de Cajá", "Suco natural de cajá", "bebida", 900},
	}

	for _, product := range products {
		if err := upsertByID(ctx, database.Collection("products"), product.id, bson.M{
			"name":        product.name,
			"description": product.description,
			"category":    product.category,
			"priceCents":  product.priceCents,
			"active":      true,
			"updatedAt":   now,
		}, bson.M{"createdAt": now}); err != nil {
			return err
		}
	}

	if err := upsertByID(ctx, database.Collection("promotions"), PromotionID, bson.M{
		"name":            "Desconto Nordeste 10",
		"code":            PromotionCode,
		"discountPercent": int64(10),
		"active":          true,
		"updatedAt":       now,
	}, bson.M{"createdAt": now}); err != nil {
		return err
	}

	stocks := []struct {
		id        string
		productID string
		quantity  int64
	}{
		{"660000000000000000000401", ProductCuscuzID, 20},
		{"660000000000000000000402", ProductTapiocaID, 15},
		{"660000000000000000000403", ProductSucoID, 30},
	}

	for _, stock := range stocks {
		if err := upsertByID(ctx, database.Collection("stock_balances"), stock.id, bson.M{
			"unitId":    UnitID,
			"productId": stock.productID,
			"quantity":  stock.quantity,
			"active":    true,
			"updatedAt": now,
		}, bson.M{"createdAt": now}); err != nil {
			return err
		}
	}

	movements := []struct {
		id        string
		productID string
		quantity  int64
	}{
		{"660000000000000000000501", ProductCuscuzID, 20},
		{"660000000000000000000502", ProductTapiocaID, 15},
		{"660000000000000000000503", ProductSucoID, 30},
	}

	for _, movement := range movements {
		if err := upsertByID(ctx, database.Collection("stock_movements"), movement.id, bson.M{}, bson.M{
			"unitId":       UnitID,
			"productId":    movement.productID,
			"type":         "ENTRADA",
			"quantity":     movement.quantity,
			"reason":       "Seed inicial",
			"balanceAfter": movement.quantity,
			"createdAt":    now,
		}); err != nil {
			return err
		}
	}

	if err := upsertByID(ctx, database.Collection("loyalty_balances"), "660000000000000000000601", bson.M{
		"userId":    ClientID,
		"points":    int64(0),
		"active":    true,
		"updatedAt": now,
	}, bson.M{"createdAt": now}); err != nil {
		return err
	}

	return nil
}

func createIndexes(ctx context.Context, database *mongo.Database) error {
	if _, err := database.Collection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}); err != nil {
		return err
	}

	if _, err := database.Collection("promotions").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "code", Value: 1}},
		Options: options.Index().SetUnique(true),
	}); err != nil {
		return err
	}

	if _, err := database.Collection("stock_balances").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "unitId", Value: 1}, {Key: "productId", Value: 1}},
		Options: options.Index().SetUnique(true),
	}); err != nil {
		return err
	}

	if _, err := database.Collection("loyalty_balances").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "userId", Value: 1}},
		Options: options.Index().SetUnique(true),
	}); err != nil {
		return err
	}

	return nil
}

func upsertByID(ctx context.Context, collection *mongo.Collection, id string, set bson.M, setOnInsert bson.M) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{}
	if len(set) > 0 {
		update["$set"] = set
	}
	if len(setOnInsert) > 0 {
		setOnInsert["_id"] = objectID
		update["$setOnInsert"] = setOnInsert
	}

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		update,
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
