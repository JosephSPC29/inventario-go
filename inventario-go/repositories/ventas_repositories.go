package repositories

import (
	"inventario-go/models"

	"gorm.io/gorm"
)

type SalesRepository interface {
	CreateSale(sale *models.Sale) error
	GetSaleByID(id uint) (*models.Sale, error)
	GetAllSales() ([]models.Sale, error)
	GetTopSellingProducts(limit int) ([]models.SaleItem, error)
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db}
}

func (r *salesRepository) CreateSale(sale *models.Sale) error {
	return r.db.Create(sale).Error
}

func (r *salesRepository) GetSaleByID(id uint) (*models.Sale, error) {
	var sale models.Sale
	if err := r.db.Preload("Items").First(&sale, id).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *salesRepository) GetAllSales() ([]models.Sale, error) {
	var sales []models.Sale
	if err := r.db.Preload("Items").Find(&sales).Error; err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *salesRepository) GetTopSellingProducts(limit int) ([]models.SaleItem, error) {
	var topProducts []models.SaleItem
	err := r.db.Model(&models.SaleItem{}).
		Select("product_id, SUM(quantity) as quantity").
		Group("product_id").
		Order("quantity DESC").
		Limit(limit).
		Scan(&topProducts).Error
	return topProducts, err
}
