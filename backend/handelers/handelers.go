package handelers

import (
	"backend/dbase"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Start(c *gin.Context) {
	num, err := strconv.ParseInt(c.Param("num"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world", "number": num})
}

func GetEstate(c *gin.Context) {
	cursor, err := dbase.DB.Collection(dbase.EstateCollection).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var estates []dbase.Estate

	for cursor.Next(context.Background()) {
		var est dbase.Estate
		if err := cursor.Decode(&est); err != nil {
			log.Fatal(err)
		}
		estates = append(estates, est)
	}

	cursor.Close(context.Background())
	c.JSON(http.StatusOK, estates)
}

func GetEstateByID(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))

	var Est dbase.Estate
	if err := dbase.DB.Collection(dbase.EstateCollection).FindOne(context.Background(), bson.M{"_id": id}).Decode(&Est); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, Est)
}

func AddEstate(c *gin.Context) {
	var Est dbase.Estate

	if err := c.BindJSON(&Est); err != nil {
		log.Fatal(err)
	}

	result, err := dbase.DB.Collection(dbase.EstateCollection).InsertOne(context.Background(), Est)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateEstate(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))
	var Est dbase.Estate

	if err := c.BindJSON(&Est); err != nil {
		log.Fatal(err)
		return
	}

	update := bson.M{"$set": Est}
	filter := bson.M{"_id": id}
	_, err := dbase.DB.Collection(dbase.EstateCollection).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estate updated successfully"})
}

func DeleteEstate(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("_id"))

	filter := bson.M{"_id": id}

	_, err := dbase.DB.Collection(dbase.EstateCollection).DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estate deleted successfully"})
}

func GetEstatesByOwnerID(c *gin.Context) {
	ownerID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner ID"})
		return
	}

	cursor, err := dbase.DB.Collection(dbase.EstateCollection).Find(context.Background(), bson.M{"OwnerID": ownerID})
	if err != nil {
		log.Fatal(err)
	}

	var estates []dbase.Estate

	for cursor.Next(context.Background()) {
		var estate dbase.Estate
		if err := cursor.Decode(&estate); err != nil {
			log.Fatal(err)
		}
		estates = append(estates, estate)
	}

	c.IndentedJSON(http.StatusOK, estates)
}

func SearchEstate(c *gin.Context) {

	fmt.Print("Im called")
}
