package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/invoice"
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
)

//mysqlInvoice used for work with invoice | Postgres
type mysqlInvoice struct {
	db *sql.DB
	storageHeader invoiceheader.Storage
	storageItems invoiceitem.Storage
}

//NewPsqlInvoice return a new pointer of mysqlInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *mysqlInvoice {
	return &mysqlInvoice{
		db: db,
		storageHeader: h,
		storageItems: i,
	}
}

//Create implement the interface of invoice.storage
func (p *mysqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil{
		tx.Rollback()
		return err
	}
	fmt.Printf("Invoice has been created, ID: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Invoice items have been created, Total %d \n", len(m.Items))

	return tx.Commit()
}

