package invoice

import (
	"github.com/cmorales95/golang_db/pkg/invoiceheader"
	"github.com/cmorales95/golang_db/pkg/invoiceitem"
)

//Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items invoiceitem.Models
}

//Storage interface that must implement a db storage
type Storage interface {
	Create(m *Model) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

//Create a new invoice
func (s *Service) Create (m *Model) error {
	return s.storage.Create(m)
}