package products

import (
	"gorm.io/gorm"
)

type Service struct {
	productRepository ProductRepository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		productRepository: *NewProductRepository(db),
	}
}

func (s *Service) GetProducts() ([]Product, error) {
	return s.productRepository.GetAll()
}

func (s *Service) FindProduct(id string) (Product, error) {
	return s.productRepository.Find(id)
}

func (s *Service) CreateProduct(name string) (Product, error) {
	p := NewProduct(name)

	err := s.productRepository.Create(p)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func (s *Service) UpdateProduct(id string, name string) error {
	p, err := s.productRepository.Find(id)
	if err != nil {
		return err
	}

	p.Name = name

	err = s.productRepository.Update(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteProduct(id string) error {
	p, err := s.productRepository.Find(id)
	if err != nil {
		return err
	}

	return s.productRepository.Delete(p)
}
