package main

import (
	"backend/handelers"
	"backend/helpers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Real estate module

func main() {
	router := gin.Default()

	godotenv.Load(".env")

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.POST("/signup/", handelers.SingUp)
	router.POST("/login/", handelers.LogIn)

	router.GET("/profile/:id", handelers.ReadProfile)
	router.PUT("/profile/:id", handelers.UpdateProfile)
	router.DELETE("/profile/:id", handelers.DelProfile)

	router.POST("/note/", handelers.CreateNote)
	router.GET("/note/", handelers.Note)
	router.GET("/note/:id", handelers.ReadNote)
	router.PUT("/note/:id", handelers.UpdateNote)
	router.DELETE("/note/:id", handelers.DelNote)

	PORT := os.Getenv("PORT")

	helpers.StartNotification()
	//helpers.EmailVerification("Mohammed", os.Getenv("MyEmail"), "google.com")

	err := router.Run("localhost:" + PORT)

	if err != nil {
		helpers.TerminationNotification(err)
	}
}
