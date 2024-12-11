package repositories

import (
	"inventario-go/models"
	"time"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetInventoryByProductID(productID uint) (*models.Inventory, error)
	UpdateInventory(inventory *models.Inventory) error
	GetLowStockProducts() ([]models.Inventory, error)
	GetExpiredProducts(currentDate time.Time) ([]models.Inventory, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) GetInventoryByProductID(productID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	if err := r.db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepository) UpdateInventory(inventory *models.Inventory) error {
	return r.db.Save(inventory).Error
}

func (r *inventoryRepository) GetLowStockProducts() ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := r.db.Where("stock < minimum_stock").Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

func (r *inventoryRepository) GetExpiredProducts(currentDate time.Time) ([]models.Inventory, error) {
	var inventories []models.Inventory
	if err := r.db.Where("expiration_date < ?", currentDate).Find(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}
