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

	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Delete: %v\n", err)
	}

}
