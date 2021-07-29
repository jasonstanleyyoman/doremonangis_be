package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/doremonangis_be/middleware"
	"github.com/jasonstanleyyoman/doremonangis_be/routes"
)

func main() {
	g := gin.Default()
	g.Use(middleware.Cors()...)
	route := routes.NewRouter()
	route.InitRouter(g)

	g.Run()
}
