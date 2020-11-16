package invoiceitem

import (
	"database/sql"
	"time"
)

//Model of invoice item
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	ProductID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(tx *sql.Tx, id uint, models Models) error
}

//Service product
type Service struct {
	Storage Storage
}

//NewService construtor return pointer of service
func NewService(s Storage) *Service {
	return &Service{Storage: s}
}

//Migrate is used for the migration of product
func (s *Service) Migrate() error {
	return s.Storage.Migrate()
}