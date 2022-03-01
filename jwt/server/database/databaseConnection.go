package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Kr-Harshit/jwt/helpers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	helpers.ErrorMsgLog("error loading .env file", err)
}

func DBinstance() *mongo.Client {
	connectionString := os.Getenv("DB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // create context with timeout enabled
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions) // creating client instance
	helpers.ErrorLog(err)

	err = client.Connect(ctx) // connecting to database
	helpers.ErrorLog(err)

	err = client.Ping(context.TODO(), nil) // checking database
	helpers.ErrorLog(err)

	fmt.Printf("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")
	var collection *mongo.Collection = client.Database(dbName).Collection(collName)
	return collection
}
