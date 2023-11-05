package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Real estate module

type estate struct {
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

const {
	Connectionstring = "mongodb+srv://moh:<password>@cluster0.7jqhbon.mongodb.net/?retryWrites=true&w=majority"
	DbName = "yt"
	CollectionName = "estate"
}

var db *mongo.Database

func init(){
	clinetOptions := options.Client().ApplytURI(ConnectionString)
	client,err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DbName)
}

func main(){
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api/estate",getEstate)
	r.GET("/api/estate/:id",getEstateById)
	r.POST("/api/estate",addEstate)
	r.PUT("/api/estate/:id", updateEstate)
	r.DELETE("/apt/estate/:id", deleteEstate)

	r.Run(":8000")
}


func getEstate(c *gin.context){

	cursor, err = db.Collection("estate").Find(context.Background(),bson.D())

	if err != nil [
		log.Fatal(err)
	]

	var estate []estate

	for cursor.Next(context.Background()){
		var employee employee
		if err := cursor.Decode(&estate); err != nil {
			log.Fatal(err)
		}
		estates = append(estates,estate)
	}

	cursor.Close(context.Background())
	c.JSON(http.StatusOK, employees)
}
