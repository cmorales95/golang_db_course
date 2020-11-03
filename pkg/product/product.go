package product

import "time"

// Model of product
type Model struct {
	ID uint
	Name string
	Observations string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Models slice of model
type Models []*Model


type Storage interface {
	Migrate() error
	/*Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error*/
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