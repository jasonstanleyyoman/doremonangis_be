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
	}
}
