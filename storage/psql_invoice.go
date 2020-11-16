package storage

import (
	"database/sql"
	"fmt"
	"github.com/cmorales95/golang_db/pkg/invoice"
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
)

//PsqlInvoice used for work with invoice | Postgres
type PsqlInvoice struct {
	db *sql.DB
	storageHeader invoiceheader.Storage
	storageItems invoiceitem.Storage
}

//NewPsqlInvoice return a new pointer of PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db: db,
		storageHeader: h,
		storageItems: i,
	}
}

//Create implement the interface of invoice.storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
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

