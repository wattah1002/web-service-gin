package router

import (
	"web-service-gin/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(dbConn *gorm.DB) {
	albumHandler := controllers.AlbumHandler{
		Db: dbConn,
	}
	router := gin.Default()
	router.GET("/albums", albumHandler.GetAlbums)
	// router.GET("/albums/:id", controllers.GetAlbumByID)
	// router.POST("/albums", controllers.PostAlbums)

	router.Run("localhost:8080")
}
