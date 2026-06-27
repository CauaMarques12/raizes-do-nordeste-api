package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func InitConnection() (*mongo.Client, *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	databaseName := os.Getenv("MONGODB_DATABASE")
	if databaseName == "" {
		databaseName = "raizes_do_nordeste"
	}

	mongoClient, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, nil); err != nil {
		panic(err)
	}

	client = mongoClient
	database = mongoClient.Database(databaseName)

	return client, database
}

func GetDatabase() *mongo.Database {
	return database
}

func GetCollection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}
