package invoiceitem

import "time"

//Model of invoice item
type Model struct {
	Id uint
	InvoiceHeaderID uint
	ProductId uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
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