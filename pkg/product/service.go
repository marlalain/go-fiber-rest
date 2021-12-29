package product

import "go-fiber-rest/pkg/entities"

type Service interface {
	Insert(product *entities.Product) (*entities.Product, error)
	Fetch() (*[]entities.Product, error)
	FindOne(ID string) (*entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Remove(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) FindOne(ID string) (*entities.Product, error) {
	return s.repository.FindOne(ID)
}

func (s service) Insert(product *entities.Product) (*entities.Product, error) {
	return s.repository.Create(product)
}

func (s service) Fetch() (*[]entities.Product, error) {
	return s.repository.Read()
}

func (s service) Update(product *entities.Product) (*entities.Product, error) {
	return s.repository.Update(product)
}

func (s service) Remove(ID string) error {
	return s.repository.Delete(ID)
}
