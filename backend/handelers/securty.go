package handelers

import (
	"backend/dbase"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {
	var content struct {
		Email        string
		Password     string
		Newpassword1 string
		Newpassword2 string
	}

	if err := c.ShouldBind(&content); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {

		var user dbase.User
		err := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), bson.M{"email": content.Email}).Decode(&user)

		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, gin.M{"Error": err.Error()})
			return
		} else {
			if content.Newpassword1 == content.Newpassword2 {
				c.IndentedJSON(http.StatusBadRequest, bson.M{"Error": "Re-enter your new password, make it identical."})
				return
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(content.Password))
				if err != nil {
					c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
					return
				} else {
					hash, err := bcrypt.GenerateFromPassword([]byte(content.Newpassword1), 10)

					if err != nil {
						c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
						return
					} else {
						result, err := dbase.DB.Collection(dbase.UserCollection).UpdateOne(
							context.Background(),
							bson.M{"email": content.Email},
							bson.M{"$set": bson.M{"password": string(hash)}},
						)
						if err != nil {
							c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
							return
						} else {
							c.IndentedJSON(http.StatusAccepted, gin.H{"result": result,
								"message": fmt.Sprintf("password has been changed to %v", content.Newpassword1)})
							return
						}
					}
				}
			}
		}

	}
}
