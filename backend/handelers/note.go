package handelers

import (
	"backend/dbase"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(c *gin.Context) {

	var note dbase.Note

	if err := c.Bind(&note); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		note.CreatedAt = time.Now()
		note.UpdatedAt = time.Now()
		result, err := dbase.DB.Collection(dbase.NoteCollection).InsertOne(context.Background(), note)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusCreated, result)
			return
		}
	}
}

func Note(c *gin.Context) {

	var notes []dbase.Note

	result, err := dbase.DB.Collection(dbase.NoteCollection).Find(context.Background(), bson.D{})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else if err := result.All(context.Background(), &notes); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	} else {
		c.IndentedJSON(http.StatusOK, notes)
		return
	}
}

func ReadNote(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		var note dbase.Note
		result := dbase.DB.Collection(dbase.NoteCollection).FindOne(context.Background(), bson.M{"_id": id})
		if err := result.Decode(&note); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusOK, note)
			return
		}
	}
}

func UpdateNote(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	var note dbase.Note

	if err := c.Bind(&note); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result, err := dbase.DB.Collection(dbase.NoteCollection).UpdateOne(context.Background(), bson.M{"_id": id}, note)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, result)
}

func DelNote(c *gin.Context) {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	result, err := dbase.DB.Collection(dbase.NoteCollection).DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}
