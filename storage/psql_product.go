package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/product"
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
	)`
	psqlCreateProduct = `INSERT INTO products (name, observations, price, created_at)
	VALUES ($1, $2, $3, $4) RETURNING id`

	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE ID = $5`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

// PsqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate .....
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration has been executed successfully")
	return nil
}

//Create implement the interface of product.storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil{
		return err
	}
	defer stmt.Close()
	
	err = stmt.QueryRow(m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return nil
	}

	fmt.Println("Se creó el producto correctamente")
	return nil
}

//GetAll implement interface of product.storage
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

//GetById implment the interface of product.storage
func (p *PsqlProduct) GetById(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()
	
	return scanRowProduct(stmt.QueryRow(id))
}

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
		timeToNull(m.UpdatedAt),
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
		return fmt.Errorf("product ID %d does not exists", m.ID)
	}

	fmt.Println("product has been updated")
	return nil
}

//Delete implement the interface of product.storage
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

	fmt.Println("the product has been deleted")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}
	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}

