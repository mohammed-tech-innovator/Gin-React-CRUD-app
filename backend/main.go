package main

import (
	"backend/handlers"
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

	router.POST("/signup/", handlers.SingUp)
	router.POST("/login/", handlers.LogIn)

	router.GET("/profile/:id", handlers.ReadProfile)
	router.PUT("/profile/:id", handlers.UpdateProfile)
	router.DELETE("/profile/:id", handlers.DelProfile)

	router.POST("/note/", handlers.CreateNote)
	router.GET("/note/", handlers.Note)
	router.GET("/note/:id", handlers.ReadNote)
	router.PUT("/note/:id", handlers.UpdateNote)
	router.DELETE("/note/:id", handlers.DelNote)

	router.GET("/verify-email/:hash/:email", handlers.VerifyEmail)
	router.PUT("/change-password/", handlers.ChangePassword)
	router.POST("/recover/", handlers.RecoverPassword)
	router.GET("/recovery/:hash/:email", func(ctx *gin.Context) {})

	PORT := os.Getenv("PORT")

	helpers.StartNotification()

	err := router.Run("localhost:" + PORT)

	if err != nil {
		helpers.TerminationNotification(err)
	}
}
