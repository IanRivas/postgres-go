package storage

import (
	"database/sql"
	"fmt"

	"github.com/IanRivas/postgres-go/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	);`
	psqlCreateProduct = `INSERT INTO products(name,observations,price,created_at) 
	VALUES($1,$2,$3,$4) RETURNING id;`
	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at 
	FROM products`
	psqlGetProductById = psqlGetAllProduct + ` WHERE id = $1`
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	price = $3, updated_at = $4 WHERE id = $5 ;`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1 ;`
)

// RETURNING para que el sql nos devuelva algo, porque el exec de postgres no devuelve nada

// para trabajar con postgres y el product
type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("entre a migrate")
		return err
	}

	fmt.Println("Migracion de producto ejecutada correctamente")
	return nil
}

// Create implement the interface product.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")
	return nil
}

// GetAll implements the interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetById implement the interface product.Storage
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement the interface product.Storage
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdateAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No existe el product con id: %d", m.ID)
	}
	fmt.Println("Se actualizo el producto correctamente")
	return nil
}

// Delete implement the interface product.Storage
func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Se elimino el producto correctamente")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updateAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdateAt = updateAtNull.Time

	return m, nil
}
