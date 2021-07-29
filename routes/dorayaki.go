package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/doremonangis_be/controller"
)

func DorayakiRoutes(g *gin.RouterGroup) {
	dorayakiController := controller.NewDorayakiController()
	dorayakiGroup := g.Group("/dorayaki")
	{
		dorayakiGroup.GET("/all", dorayakiController.GetAllDorayaki())
		dorayakiGroup.POST("/new", dorayakiController.AddDorayaki())
		dorayakiGroup.DELETE("/:id", dorayakiController.DeleteDorayaki())
		dorayakiGroup.GET("/:id", dorayakiController.GetDorayakiInfo())
	}
}
