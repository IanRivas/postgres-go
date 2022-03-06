package storage

import (
	"database/sql"
	"fmt"
)

//le dice migrate a CREAR
const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
		(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE
		RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY
		(product_id) REFERENCES products (id) ON UPDATE
		RESTRICT ON DELETE RESTRICT
	);`
)

// para trabajar con postgres y el product
type PsqlInvoiceItem struct {
	db *sql.DB
}

func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implement the interface invoiceItem.Storage
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	fmt.Println(stmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("entre a migrate")
		return err
	}

	fmt.Println("Migracion de invoiceItem ejecutada correctamente")
	return nil
}
