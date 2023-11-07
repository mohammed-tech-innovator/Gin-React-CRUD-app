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

	router.GET("/:num", handelers.Start)
	router.GET("/estate", handelers.GetEstate)
	router.GET("/estate/:id", handelers.GetEstateByID)
	router.POST("/estate", handelers.AddEstate)
	router.PUT("/estate/:id", handelers.UpdateEstate)
	router.DELETE("/estate/:id", handelers.DeleteEstate)
	router.GET("/owner/:id", handelers.GetOwnerByID)
	router.GET("/owners/", handelers.GetOwners)
	router.GET("/estatebyownerid/:id", handelers.GetEstatesByOwnerID)
	router.POST("/search/", handelers.SearchEstate)

	router.Run("localhost:8000")
}
