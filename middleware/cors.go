package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func header(g *gin.Context) {
	g.Header("Access-Control-Allow-Origin", "*")
	g.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
}

func Cors() gin.HandlersChain {

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Accept"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
	}

	return []gin.HandlerFunc{header, cors.New(config)}
}
