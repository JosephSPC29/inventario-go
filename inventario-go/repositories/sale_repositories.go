package repositories

import (
	"inventario-go/models"

	"gorm.io/gorm"
)

type SaleRepository interface {
	CreateSale(sale *models.Sale) error
	GetSales() ([]models.Sale, error)
	GetSaleByID(id uint) (*models.Sale, error)
	UpdateSale(sale *models.Sale) error
	DeleteSale(id uint) error
}

type saleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &saleRepository{db}
}

func (r *saleRepository) CreateSale(sale *models.Sale) error {
	if err := r.db.Create(sale).Error; err != nil {
		return err
	}
	return nil
}

func (r *saleRepository) GetSales() ([]models.Sale, error) {
	var sales []models.Sale
	if err := r.db.Find(&sales).Error; err != nil {
		return nil, err
	}
	return sales, nil
}

func (r *saleRepository) GetSaleByID(id uint) (*models.Sale, error) {
	var sale models.Sale
	if err := r.db.First(&sale, id).Error; err != nil {
		return nil, err
	}
	return &sale, nil
}

func (r *saleRepository) UpdateSale(sale *models.Sale) error {
	if err := r.db.Save(sale).Error; err != nil {
		return err
	}
	return nil
}

func (r *saleRepository) DeleteSale(id uint) error {
	if err := r.db.Delete(&models.Sale{}, id).Error; err != nil {
		return err
	}
	return nil
}
