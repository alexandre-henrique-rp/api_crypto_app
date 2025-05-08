package models

import "time"

type StockExchange struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Price float64 `json:"price"`
	Valor float64 `json:"valor"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	TableName = "stock_exchange"
	ComandSQL = `CREATE TABLE IF NOT EXISTS stock_exchange (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		symbol TEXT NOT NULL,
		price REAL NOT NULL,
		valor REAL NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
)