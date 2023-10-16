package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/controllers"
)

func main() {
	router := gin.Default()
	v1_group := router.Group("/v1")
	{
		v1_group.GET("/albums", controllers.GetAlbums)
		v1_group.GET("/albums/:id", controllers.GetAlbumByID)
		v1_group.POST("/albums", controllers.PostAlbums)
	}

	router.Run("localhost:8080")
}
