package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registra as rotas da API
func RegisterRoutes(router *gin.Engine) {
	app := router.Group("/app")
	// Exemplo de rota GET
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}