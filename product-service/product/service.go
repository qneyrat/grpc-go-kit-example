package product

import "workshop-go-kit/product-service/database"

type Service interface {
	GetProduct(id int32) (*database.Product, error)
	CreateProduct(name string) (*database.Product, error)
}

type ProductService struct {
	Repository database.ProductRepository
}

func (s *ProductService) GetProduct(id int32) (*database.Product, error) {
	return s.Repository.FindById(id)
}

func (s *ProductService) CreateProduct(name string) (*database.Product, error) {
	return s.Repository.Create(name)
}
