package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes registra as rotas da API
func RegisterRoutes(router *gin.Engine) {
	// Exemplo de rota GET
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}