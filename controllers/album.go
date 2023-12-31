package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"web-service-gin/models"
)

type AlbumHandler struct {
	Db *gorm.DB
}

// albums slice to seed record album data.
// var albums = []models.Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// getAlbums responds with the list of all albums as JSON.
func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	var albums []models.Album
	h.Db.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

// // postAlbums adds an album from JSON received in the request body.
func (h *AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	h.Db.Create(&newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
// func GetAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
