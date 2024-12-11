package repositories

import (
	"inventario-go/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*models.Product, error)
	GetProducts(filter map[string]interface{}) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetProducts(filter map[string]interface{}) ([]models.Product, error) {
	var products []models.Product
	query := r.db

	if name, ok := filter["name"].(string); ok {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if category, ok := filter["category"].(string); ok {
		query = query.Where("category = ?", category)
	}
	if minPrice, ok := filter["min_price"].(float64); ok {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice, ok := filter["max_price"].(float64); ok {
		query = query.Where("price <= ?", maxPrice)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
