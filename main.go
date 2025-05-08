package main

import (
	"fmt"
	"log"

	"api_crypto_app/config"
	"api_crypto_app/models"
	"api_crypto_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := config.Connect()
	fmt.Println(DB)
	fmt.Println("Conectado a database")
	_, errModel := DB.Exec(models.ComandSQL)
	if errModel != nil {
		log.Fatalf("Erro ao criar tabela: %v", errModel)
	}
	fmt.Println("Tabela criada com sucesso")
	defer DB.Close()
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
