package handelers

import (
	"backend/dbase"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadProfile(c *gin.Context) {
	id := c.Param("id")

	var user dbase.User

	result := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), bson.D{{"_id", id}})

	if err := result.Decode(&user); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, user)

}

func UpdateProfile(c *gin.Context) {
	id := c.Param("id")

	var user dbase.User

	if err := c.Bind(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result, err := dbase.DB.Collection(dbase.UserCollection).UpdateOne(context.Background(), bson.D{{"_id", id}}, user)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

	c.IndentedJSON(http.StatusAccepted, result)
}

func DelProfile(c *gin.Context) {
	id := c.Param("id")

	result, err := dbase.DB.Collection(dbase.UserCollection).DeleteOne(context.Background(), bson.D{{"_id", id}})

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, result)

}
