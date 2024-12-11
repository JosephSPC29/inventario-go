package services

import (
	"errors"
	"time"

	"inventario-go/models"
	"inventario-go/repositories"
)

type InventoryService interface {
	RegisterStock(productID uint, stock int, minimumStock int, expirationDate *time.Time) error
	UpdateStock(productID uint, stockChange int) error
	GetLowStockAlerts() ([]models.Inventory, error)
	GetExpiredProducts() ([]models.Inventory, error)
}

type inventoryService struct {
	repo repositories.InventoryRepository
}

func NewInventoryService(repo repositories.InventoryRepository) InventoryService {
	return &inventoryService{repo}
}

func (s *inventoryService) RegisterStock(productID uint, stock int, minimumStock int, expirationDate *time.Time) error {
	inventory, err := s.repo.GetInventoryByProductID(productID)
	if err == nil && inventory != nil {
		return errors.New("inventory already exists for this product")
	}

	newInventory := &models.Inventory{
		ProductoID:      productID,
		Stock:           stock,
		StockMinimo:     minimumStock,
		FechaExpiracion: expirationDate,
	}
	return s.repo.UpdateInventory(newInventory)
}

func (s *inventoryService) UpdateStock(productID uint, stockChange int) error {
	inventory, err := s.repo.GetInventoryByProductID(productID)
	if err != nil {
		return errors.New("product not found in inventory")
	}

	inventory.Stock += stockChange
	if inventory.Stock < 0 {
		return errors.New("insufficient stock")
	}
	return s.repo.UpdateInventory(inventory)
}

func (s *inventoryService) GetLowStockAlerts() ([]models.Inventory, error) {
	return s.repo.GetLowStockProducts()
}

func (s *inventoryService) GetExpiredProducts() ([]models.Inventory, error) {
	return s.repo.GetExpiredProducts(time.Now())
}
