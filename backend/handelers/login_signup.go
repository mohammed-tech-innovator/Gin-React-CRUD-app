package handelers

import (
	"backend/dbase"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	godotenv.Load("./.env")
}

func SingUp(c *gin.Context) {

	var content struct {
		Fname    string
		Lname    string
		Password string
		Email    string
	}

	if c.Bind(&content) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Failed to match input data."})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(content.Password), 10)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	user := dbase.User{Email: content.Email, Fname: content.Fname,
		Lname: content.Lname, Password: string(hash), CreatedAt: time.Now()}

	result, err := dbase.DB.Collection(dbase.UserCollection).InsertOne(context.Background(), user)
	if err != nil {
		c.IndentedJSON(http.StatusNotImplemented, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusCreated, result)

}

func LogIn(c *gin.Context) {

	var content struct {
		Email    string
		Password string
	}

	var user dbase.User

	if c.Bind(&content) != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Failed to match input data."})
		return
	}

	result := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), gin.H{"Email": content.Email})
	if err := result.Decode(&user); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(content.Password))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Wrong password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
