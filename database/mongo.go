package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnection() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")

	}

	url := os.Getenv("MONGODB_CONNECTION")

	client, err := mongo.NewClient(options.Client().ApplyURI((url)))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to Database")

	return client
}

var client *mongo.Client = DbConnection()

// Function to OpenCollection in Database
func openCollection(dbClient *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = dbClient.Database("store-management").Collection(collectionName)
	return collection
}

// Collections Variables - Mudasir Ali
var InventoryCollection *mongo.Collection = openCollection(client, "inventory")
