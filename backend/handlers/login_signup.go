package handlers

import (
	"backend/dbase"
	"backend/helpers"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	godotenv.Load("./.env")
}

func SingUp(c *gin.Context) {

	var content struct {
		Fname     string
		Lname     string
		Email     string
		Password  string
		RPassword string
	}

	if err := c.ShouldBind(&content); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "tag": "failed to fetch data"})
		return
	} else {

		if content.Password != content.RPassword {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Password mismatch", "tag": "⚠️ Please Enter the same Password 🔤"})
			return
		} else if helpers.IsNumeric(content.Password) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Password isn't valid", "tag": "⚠️Password is entirly numeric 🔢"})
			return
		} else if len(content.Password) < 6 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "short password", "tag": "⚠️ The password should be at least 6 characters 🔤"})
			return
		} else if len(content.Fname) < 1 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "No frist name", "tag": "⚠️ please enter your first name"})
			return
		} else if len(content.Lname) < 1 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "No last name", "tag": "⚠️ please enter your last name"})
			return
		} else {

			cout, err := dbase.DB.Collection(dbase.UserCollection).CountDocuments(context.Background(), bson.M{"email": content.Email})
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "tag": "⚠️ database Error please try again 🤦‍♂️."})
				return
			} else if cout > 0 {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "This Email is already used.",
					"tag": "⚠️ This Email is already associated with an accout please login or enter a new email📧."})
				return
			} else {
				hash, err := bcrypt.GenerateFromPassword([]byte(content.Password), 10)

				if err != nil {
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "tag": "⚠️faild to process password."})
					return
				} else {

					var user dbase.User = dbase.User{
						Fname:         content.Fname,
						Lname:         content.Lname,
						Email:         content.Email,
						Password:      string(hash),
						CreatedAt:     time.Now(),
						EmailVerified: false,
					}

					result, err := dbase.DB.Collection(dbase.UserCollection).InsertOne(context.Background(), user)
					if err != nil {
						c.IndentedJSON(http.StatusNotImplemented, gin.H{"err": err.Error(),
							"tag": "⚠️database Error please try again 🤦‍♂️."})
						return
					} else {
						emailHashed := sha256.Sum256([]byte(fmt.Sprintf("%s%s", os.Getenv("VERKEY"), user.Email)))
						encodedHash := base64.URLEncoding.EncodeToString(emailHashed[:])
						url := fmt.Sprintf("%sverify-email/%v/%v", os.Getenv("ROOTURL"), encodedHash, user.Email)
						go helpers.EmailVerification(fmt.Sprintf(user.Fname+" "+user.Lname), user.Email,
							url)

						token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
							"sub": user.ID,
							"exp": time.Now().Add(time.Hour).Unix(),
						})

						tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

						if err != nil {
							c.JSON(http.StatusInternalServerError, gin.H{
								"Error": err.Error(),
							})
							return

						} else {

							c.JSON(http.StatusOK, gin.H{
								"token":  tokenString,
								"result": result,
							})
							return
						}
					}

				}
			}
		}
	}
}

func LogIn(c *gin.Context) {

	var content struct {
		Email    string
		Password string
	}

	var user dbase.User

	if err := c.ShouldBind(&content); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "tag": "Bad Data"})
		return
	} else {

		result := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), gin.H{"email": content.Email})
		if err := result.Decode(&user); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error(), "tag": "No Such User"})
			return
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(content.Password))
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "tag": "Wrong Password"})
				return
			} else {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"sub": user.ID,
					"exp": time.Now().Add(time.Hour).Unix(),
				})

				tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

				if err != nil {
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "tag": "Couldn't Process Password"})
				} else {
					c.IndentedJSON(http.StatusOK, gin.H{
						"token": tokenString,
					})
				}
			}
		}
	}
}

func VerifyEmail(c *gin.Context) {

	hash := c.Param("hash")
	email := c.Param("email")

	emailHashed := sha256.Sum256([]byte(fmt.Sprintf("%s%s", os.Getenv("VERKEY"), email)))
	encodedHash := base64.URLEncoding.EncodeToString(emailHashed[:])

	if subtle.ConstantTimeCompare([]byte(hash), []byte(encodedHash)) == 1 {

		result, err := dbase.DB.Collection(dbase.UserCollection).UpdateOne(
			context.Background(),
			bson.M{"email": email},
			bson.M{"$set": bson.M{"emailverified": true}},
		)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusAccepted, gin.H{"result": result})
		}

	} else {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"Error": fmt.Sprintf("Your access has been denaied%v : %v", hash, emailHashed)})
	}

}
