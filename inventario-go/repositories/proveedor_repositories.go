package repositories

import (
	"inventario-go/models"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	CreateSupplier(supplier *models.Supplier) error
	UpdateSupplier(supplier *models.Supplier) error
	DeleteSupplier(id uint) error
	GetSupplierByID(id uint) (*models.Supplier, error)
	GetAllSuppliers() ([]models.Supplier, error)
	AddSupplierHistory(history *models.SupplierHistory) error
	GetSupplierHistory(supplierID uint) ([]models.SupplierHistory, error)
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) CreateSupplier(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) UpdateSupplier(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) DeleteSupplier(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}

func (r *supplierRepository) GetSupplierByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	if err := r.db.First(&supplier, id).Error; err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierRepository) GetAllSuppliers() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	if err := r.db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (r *supplierRepository) AddSupplierHistory(history *models.SupplierHistory) error {
	return r.db.Create(history).Error
}

func (r *supplierRepository) GetSupplierHistory(supplierID uint) ([]models.SupplierHistory, error) {
	var history []models.SupplierHistory
	if err := r.db.Where("supplier_id = ?", supplierID).Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}
