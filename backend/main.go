package main

import (
	"backend/handelers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Real estate module

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.GET("/:num", handelers.Hello)
	router.GET("/api/estate", handelers.GetEstate)
	router.GET("/api/estate/:id", handelers.GetEstateByID)
	router.POST("/api/estate", handelers.AddEstate)
	router.PUT("/api/estate/:id", handelers.UpdateEstate)
	router.DELETE("/api/estate/:id", handelers.DeleteEstate)

	router.Run("localhost:8000")
}
