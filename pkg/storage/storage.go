package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// singleton para open
func NewPostgresDB() {
	// lo que esta en la funcion anonima ,se va a ejecutar una sola vez , aunque ejecutemos NewPostgresDB varias veces
	once.Do(func() {
		var err error
		connStr := "postgres://postgres:nobunaga@localhost:5432/golang?sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Cant open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Cant do ping: %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}