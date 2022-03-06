package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrIDNotFound = errors.New("El producto no contiene un ID")
)

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdateAt     time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s\n",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdateAt.Format("2006-01-02"))
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}

// Service of product
type Service struct {
	storage Storage
}

// return a pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create is used for create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used for get all the products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetById is used for get a product
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Update is used for update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdateAt = time.Now()

	return s.storage.Update(m)
}

// Delete is used for delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
