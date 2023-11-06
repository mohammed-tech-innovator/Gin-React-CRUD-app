package dbase

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ConnectionString = "mongodb+srv://moh:<>@cluster0.7jqhbon.mongodb.net/?retryWrites=true&w=majority"
	DbName           = "realstateapp"
	EstateCollection = "Estate"
	OwnerCollection  = "Owner"
)

var DB *mongo.Database

func init() {
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database(DbName)
}
