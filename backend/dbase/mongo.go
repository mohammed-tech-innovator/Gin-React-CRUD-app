package dbase

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName         = "CRUD"
	UserCollection = "User"
	NoteCollection = "Note"
	MetaCollection = "Meta"
)

var DB *mongo.Database
var ConnectionString string

func init() {

	godotenv.Load("./.env")
	ConnectionString = os.Getenv("DBCON")

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
