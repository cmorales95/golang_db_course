package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
)

const (
	mysqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
    	id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    	client VARCHAR(100) NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT now(),
    	updated_at TIMESTAMP
	)`

	mysqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

//MySQLInvoiceHeader used for work with postgres - product
type MySQLInvoiceHeader struct {
	db *sql.DB
}

//NewMySQLInvoiceHeader return a new pointer of MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

//Migrate psql invoice.header
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product invoiceHeader has been executed succesfully")
	return nil
}

//CreateTx implement interface of invoiceheader.storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx,m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()


	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)
	return nil
}

