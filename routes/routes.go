package routes

import (
	"github.com/ChristinaKozi/go-gin-gonic/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	songRoutes := router.Group("/songs")
	{
		songRoutes.POST("/new", controllers.CreateSong)
		songRoutes.GET("/", controllers.GetSongs)
		songRoutes.GET("/:id", controllers.GetSongByID)
		songRoutes.PUT("/:id", controllers.UpdateSong)
		songRoutes.DELETE("/:id", controllers.DeleteSong)
	}
}
