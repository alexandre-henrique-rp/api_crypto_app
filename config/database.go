package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

// Connect inicializa a conexão com o banco de dados SQLite
func Connect() {
	var err error
	DB, err = sql.Open("sqlite3", "./crypto.db")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Testa a conexão
	if err := DB.Ping(); err != nil {
		log.Fatalf("Erro ao testar conexão com banco de dados: %v", err)
	}
	criarTabelaCryptos()
}

// criarTabelaCryptos cria a tabela de criptomoedas se ela não existir
func criarTabelaCryptos() {
	sqlCreateTable := `
	CREATE TABLE IF NOT EXISTS cryptos (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		nome TEXT NOT NULL,
		tag TEXT NOT NULL,
		price_usd TEXT NOT NULL,
		price_br TEXT NOT NULL,
		porcentagem_hora TEXT NOT NULL,
		porcentagem_dia TEXT NOT NULL,
		porcentagem_semana TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(sqlCreateTable)
	if err != nil {
		log.Fatalf("Erro ao criar tabela cryptos: %v", err)
	}
}