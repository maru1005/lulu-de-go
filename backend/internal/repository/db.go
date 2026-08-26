package repository

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("DBへのPingに失敗しました: %v", err)
	}

	return db
}
