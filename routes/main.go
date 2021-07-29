package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (router *Router) InitRouter(g *gin.Engine) {
	versionGroup := g.Group("/v1")
	StoreRoute(versionGroup)
	DorayakiRoutes(versionGroup)
}

func NewRouter() Router {
	return Router{}
}
