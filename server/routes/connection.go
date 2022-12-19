package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//return type struct of client which was present in mongodb
func DBinstance() *mongo.Client {
	MongoDb := "URL"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb)) //connecting
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //timeout
	defer cancel()
	err = client.Connect(ctx) //connecting to db
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

//opening calori collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("caloriesdb").Collection(collectionName)
	return collection
}
