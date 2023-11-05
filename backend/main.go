package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Real estate module

type Estate struct {
	ID        primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	TITLE     string                 `json:"title,omitempty" bson:"title,omitempty"`
	OwnerID   primitive.ObjectID     `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
	Latitude  float64                `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude float64                `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Price     float64                `json:"price,omitempty" bson:"price,omitempty"`
	Currency  string                 `json:"currency,omitempty" bson:"currency,omitempty"`
	Photos    []string               `json:"photos,omitempty" bson:"photos,omitempty"`
	Features  map[string]interface{} `json:"features,omitempty" bson:"features,omitempty"`
	CreatedAt time.Time              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Owner struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
}

const (
	ConnectionString = "mongodb+srv://moh:<password>@cluster0.7jqhbon.mongodb.net/?retryWrites=true&w=majority"
	DbName           = "yt"
	EstateCollection = "estate"
	OwnerCollection  = "owner"
)

var db *mongo.Database

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
	db = client.Database(DbName)
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	r.GET("/api/estate", getEstate)
	r.GET("/api/estate/:id", getEstateByID)
	r.POST("/api/estate", addEstate)
	r.PUT("/api/estate/:id", updateEstate)
	r.DELETE("/api/estate/:id", deleteEstate)

	r.Run(":8000")
}

func getEstate(c *gin.Context) {
	cursor, err := db.Collection(EstateCollection).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var estates []Estate

	for cursor.Next(context.Background()) {
		var est Estate
		if err := cursor.Decode(&est); err != nil {
			log.Fatal(err)
		}
		estates = append(estates, est)
	}

	cursor.Close(context.Background())
	c.JSON(http.StatusOK, estates)
}

func getEstateByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))

	var Est Estate
	if err := db.Collection(EstateCollection).FindOne(context.Background(), bson.M{"_id": id}).Decode(&Est); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, Est)
}

func addEstate(c *gin.Context) {
	var Est Estate

	if err := c.BindJSON(&Est); err != nil {
		log.Fatal(err)
	}

	result, err := db.Collection(EstateCollection).InsertOne(context.Background(), Est)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func updateEstate(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var Est Estate

	if err := c.BindJSON(&Est); err != nil {
		log.Fatal(err)
		return
	}

	update := bson.M{"$set": Est}
	filter := bson.M{"_id": id}
	_, err := db.Collection(EstateCollection).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estate updated successfully"})
}
