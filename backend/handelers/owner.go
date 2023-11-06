package handelers

import (
	"backend/dbase"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOwnerByID(c *gin.Context) {
	owner_id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var owner dbase.Owner

	if err := dbase.DB.Collection(dbase.OwnerCollection).FindOne(context.Background(), bson.M{"_id": owner_id}).Decode(&owner); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"status": "NotFound", "err": err})
		return
	}

	c.IndentedJSON(http.StatusOK, owner)
}

func GetOwners(c *gin.Context) {
	cursor, err := dbase.DB.Collection(dbase.OwnerCollection).Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var owners []dbase.Owner

	for cursor.Next(context.Background()) {
		var owner dbase.Owner
		if err := cursor.Decode(&owner); err != nil {
			log.Fatal(err)
		}
		owners = append(owners, owner)
	}

	cursor.Close(context.Background())
	c.IndentedJSON(http.StatusOK, owners)
}
