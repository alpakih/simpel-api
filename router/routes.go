package router

import (
	"github.com/alpakih/simpel-api/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Simpel Api Example Gin")
	})
	orderGroup := router.Group("/orders")
	orderController := controllers.OrderController(db)
	orderGroup.GET("/:id", orderController.GetByID)
	orderGroup.GET("/", orderController.GetAll)
	orderGroup.POST("/", orderController.Store)
	orderGroup.DELETE("/:id", orderController.Destroy)
	orderGroup.PUT("/", orderController.Update)
	return router
}
