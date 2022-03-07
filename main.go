package main

import (
	"log"

	"github.com/IanRivas/postgres-go/pkg/invoice"
	"github.com/IanRivas/postgres-go/pkg/invoiceheader"
	"github.com/IanRivas/postgres-go/pkg/invoiceitem"
	"github.com/IanRivas/postgres-go/pkg/storage"
	_ "github.com/lib/pq"
)

func main() {

	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "nerd profile",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
			&invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
