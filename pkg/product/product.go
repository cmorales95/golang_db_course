package product

import (
	"errors"
	"fmt"
	"time"
)

var errIdNotFound = errors.New("product Id is not setting")

// Model of product
type Model struct {
	ID uint
	Name string
	Observations string
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s\n",
		m.ID, m.Name, m.Observations, m.Price, m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"))
}

//Models slice of model
type Models []*Model


type Storage interface {
	Migrate() error
	Create(model *Model) error
	GetById(id uint) (*Model, error)
	GetAll() (Models, error)
	Update(model *Model) error
	Delete(id uint) error
}

//Service product
type Service struct {
	Storage Storage
}

//NewService constructor return pointer of service
func NewService(s Storage) *Service {
	return &Service{Storage: s}
}

//Migrate is used for the migration of product
func (s *Service) Migrate() error {
	return s.Storage.Migrate()
}

//Create product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.Storage.Create(m)
}

//GetAll is used for get all products
func (s *Service) GetAll() (Models, error) {
	return s.Storage.GetAll()
}

//GetById is used to get a specific product by id
func (s *Service) GetById(id uint) (*Model, error) {
	return s.Storage.GetById(id)
}

//Update is used to update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return errIdNotFound
	}
	m.UpdatedAt = time.Now()

	return s.Storage.Update(m)
}

//Delete is used to update a product
func (s *Service) Delete(id uint) error {
	return s.Storage.Delete(id)
}