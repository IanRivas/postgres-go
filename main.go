package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/IanRivas/postgres-go/pkg/product"
	"github.com/IanRivas/postgres-go/pkg/storage"
	_ "github.com/lib/pq"
)

func main() {

	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetByID(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con este id")
	case err != nil:
		log.Fatalf("product.GetById: %v\n", err)
	default:
		fmt.Println(m)
	}

}
