package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// Connect inicializa a conexão com o banco de dados SQLite
func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %v", err)
	}
  Db_Phat := os.Getenv("DB_HOST")


	DB, errDb := sql.Open("sqlite3", Db_Phat)
	if errDb != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", errDb)
	}

	//teste de conexão
	teste := DB.Ping()
	if teste != nil {
		log.Fatalf("Erro ao testar conexão com banco de dados: %v", teste)
	}

	return DB
}
