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

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {

	var content helpers.ChangePasswordForm

	if err := c.ShouldBind(&content); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {

		var user dbase.User
		err := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), bson.M{"email": content.Email}).Decode(&user)

		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, bson.M{"Error": err.Error()})
			return
		} else {
			if content.Newpassword1 != content.Newpassword2 {
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
							helpers.SendMailSimple("password change notification", " Your password has been successfully changed. If you initiated this action, no further steps are needed.\nIf you didn't request this change or if you have any concerns, please contact our support team immediately to ensure the security of your account.", []string{user.Email})
							c.IndentedJSON(http.StatusAccepted, gin.H{"result": result, "message": fmt.Sprintf("password has been changed to %v", content.Newpassword1)})

							return
						}
					}
				}
			}
		}
	}
}

func RecoverPassword(c *gin.Context) {

	var content helpers.RecoverPasswordForm

	if err := c.ShouldBind(&content); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(), "tag": "⚠️Couldn't fitch data."})
		return
	} else if errMessage, tagMessage, err := helpers.CheckRecoverPasswordForm(content); err {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": errMessage, "tag": tagMessage})
		return
	} else {
		result := dbase.DB.Collection(dbase.UserCollection).FindOne(context.Background(), bson.M{"email": content.Email})

		if err := result.Err(); err == mongo.ErrNoDocuments {

			c.IndentedJSON(http.StatusNotFound, gin.H{"Error": err.Error(), "tag": "⚠️No Such Email in Our Database 🛑"})

			return

		} else if err != nil {

			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "tag": "⚠️ An Error occurred ."})

			return

		} else {

			var user dbase.User

			if err := result.Decode(&user); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "tag": "⚠️ An Error occurred ."})
				return
			} else {

				if user.EmailVerified {

					emailHashed := sha256.Sum256([]byte(fmt.Sprintf("%s%s", os.Getenv("VERKEY"), user.Email)))
					encodedHash := base64.URLEncoding.EncodeToString(emailHashed[:])
					url := fmt.Sprintf("%srecovery/%v/%v/", os.Getenv("FEROOTURL"), encodedHash, user.Email)

					defer func() {
						go helpers.PWRecoveryEmail(fmt.Sprintf("%v %v", user.Fname, user.Lname), user.Email, url)
					}()

					c.IndentedJSON(http.StatusAccepted, gin.H{"tag": "Recovery Email has been sent, please check your email."})

				} else {
					c.IndentedJSON(http.StatusNotAcceptable, gin.H{"Error": "Email is not verified", "tag": "⚠️Your Email is not verified please contact us."})
				}

			}

		}
	}

}

func ChangeRecoverdPassword(c *gin.Context) {

	hash := c.Param("hash")
	email := c.Param("email")

	emailHashed := sha256.Sum256([]byte(fmt.Sprintf("%s%s", os.Getenv("VERKEY"), email)))
	encodedHash := base64.URLEncoding.EncodeToString(emailHashed[:])

	if subtle.ConstantTimeCompare([]byte(hash), []byte(encodedHash)) == 1 {

		var pwcr helpers.PWCR

		if err := c.ShouldBind(&pwcr); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error(),
				"tag": "⚠️ An error occurred while processing the request."})
			return
		} else if errMessage, tagMessage, err := helpers.CheckPWCR(pwcr); err {
			c.IndentedJSON(http.StatusBadRequest, bson.M{"Error": errMessage, "tag": tagMessage})
			return
		} else {

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwcr.NPassword1), 10)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error(), "tag": "🛑please try again."})
				return
			} else {
				result, err := dbase.DB.Collection(dbase.UserCollection).UpdateOne(context.Background(),
					bson.M{"email": email}, bson.M{"$set": bson.M{"password": string(hashedPassword)}})

				if err != nil {
					c.IndentedJSON(http.StatusNotFound, bson.M{"Error": err.Error(), "tag": "⚠️ no such user."})
					return
				} else {
					defer func() {
						go helpers.PWRCEmail(email, fmt.Sprintf("%s/%s", os.Getenv("FEROOTURL"), "login/"))
					}()
					c.IndentedJSON(http.StatusAccepted, bson.M{"result": result, "tag": "password has been updated successfuly👌."})
					return
				}
			}

		}

	} else {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"Error": "Access has been denaid", "tag": fmt.Sprintf("Your access has been denaied%v : %v", hash, emailHashed)})
		return
	}

}
