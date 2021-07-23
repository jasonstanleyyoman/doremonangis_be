package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonstanleyyoman/doremonangis_be/routes"
)

func main() {
	g := gin.Default()
	route := routes.NewRouter()
	route.InitRouter(g)

	g.Run()
}
