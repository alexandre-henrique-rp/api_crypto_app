package main

import (
	"github.com/alexandre-henrique-rp/go_api_crypto/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
