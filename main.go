package main

import (
	"log"

	"github.com/IanRivas/postgres-go/pkg/product"
	"github.com/IanRivas/postgres-go/pkg/storage"
	_ "github.com/lib/pq"
)

func main() {

	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v\n", err)
	}
}
