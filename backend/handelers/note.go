package handelers

import (
	"backend/dbase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNote(c *gin.Context) {

	var note dbase.Note

	if err := c.Bind(&note); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result, err := dbase.DB.Collection(dbase.NoteCollection).InsertOne(context.Background(), note)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func Note(c *gin.Context) {

	var notes []dbase.Note

	result, err := dbase.DB.Collection(dbase.NoteCollection).Find(context.Background(), bson.D{})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if err := result.All(context.Background(), &notes); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, notes)
}

func ReadNote(c *gin.Context) {
	id := c.Param("id")

	var note dbase.Note

	result := dbase.DB.Collection(dbase.NoteCollection).FindOne(context.Background(), gin.H{"_id": id})

	if err := result.Decode(&note); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, note)

}

func UpdateNote(c *gin.Context) {
	id := c.Param("id")

	var note dbase.Note

	if err := c.Bind(&note); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result, err := dbase.DB.Collection(dbase.NoteCollection).UpdateOne(context.Background(), gin.H{"_id": id}, note)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, result)
}

func DelNote(c *gin.Context) {
	id := c.Param("id")

	result, err := dbase.DB.Collection(dbase.NoteCollection).DeleteOne(context.Background(), gin.H{"_id": id})

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, result)

}
