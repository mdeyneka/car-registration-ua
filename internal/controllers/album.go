package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdeineka/car-registration-ua/internal/database"
	"github.com/mdeineka/car-registration-ua/internal/models"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Albums)
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	database.Albums = append(database.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range database.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
