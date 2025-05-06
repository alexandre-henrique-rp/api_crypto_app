package main

import (
	"fmt"

	database "api_crypto_app/config"
	"api_crypto_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	fmt.Println("Conectado a database")
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
