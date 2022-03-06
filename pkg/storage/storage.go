package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

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

// Helper intermedio para ver si un string es null en sql
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// Helper intermedio para ver si un timestamp es null en sql
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
