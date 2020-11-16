package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
)

const (
	mysqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_item(
		id INT AUTO_INCREMENT NOT NULL primary key ,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY(invoice_header_id) 
			REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY(product_id)
			REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`

	mysqlCreateInvoiceItem = `INSERT INTO invoice_item(invoice_header_id, product_id) VALUES (?, ?)`
)

//MySQLInvoiceItem used for work with postgres - product
type MySQLInvoiceItem struct {
	db *sql.DB
}

//NewMySQLInvoiceItem return a new pointer of MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

//Migrate psql invoice.header
func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product invoice_item has been executed succesfully")
	return nil
}

//CreateTx implements invoiceitem.createtx
func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, id uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		result, err := stmt.Exec(id, item.ProductID)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil{
			return err
		}

		item.ID = uint(id)
	}
	return nil
}
