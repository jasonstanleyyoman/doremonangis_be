package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/doremonangis_be/controller"
)

func StoreRoute(g *gin.RouterGroup) {
	storeController := controller.NewStoreController()
	storeGroup := g.Group("/store")
	{
		storeGroup.GET("/all", storeController.GetAllStore())
		storeGroup.POST("/new", storeController.AddStore())
		storeGroup.POST("/dorayaki/add", storeController.AddStock())
		storeGroup.POST("/dorayaki/remove", storeController.RemoveStock())
		storeGroup.POST("/dorayaki/move", storeController.MoveStock())
		storeGroup.DELETE("/:id", storeController.DeleteStore())
		storeGroup.GET("/:id", storeController.GetStoreInfo())
	}
}
