package services

import (
	"errors"
	"inventario-go/models"
	"inventario-go/repositories"
)

type ProductService interface {
	RegisterProduct(name, description, category string, price float64) error
	UpdateProduct(id uint, name, description, category string, price float64) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*models.Product, error)
	GetProducts(filter map[string]interface{}) ([]models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) RegisterProduct(name, description, category string, price float64) error {
	product := &models.Product{
		Nombre:      name,
		Descripcion: description,
		Categoria:   category,
		Precio:      price,
	}
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(id uint, name, description, category string, price float64) error {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	product.Nombre = name
	product.Descripcion = description
	product.Categoria = category
	product.Precio = price
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *productService) GetProducts(filter map[string]interface{}) ([]models.Product, error) {
	return s.repo.GetProducts(filter)
}
