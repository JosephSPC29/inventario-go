package services

import (
	"errors"
	"inventario-go/models"
	"inventario-go/repositories"
	"time"
)

type SupplierService interface {
	RegisterSupplier(name, contactInfo string) error
	UpdateSupplier(id uint, name, contactInfo string, rating float64) error
	DeleteSupplier(id uint) error
	GetSupplierByID(id uint) (*models.Supplier, error)
	GetAllSuppliers() ([]models.Supplier, error)
	AddSupplierHistory(supplierID uint, orderID string, amount float64, description string) error
	GetSupplierHistory(supplierID uint) ([]models.SupplierHistory, error)
}

type supplierService struct {
	repo repositories.SupplierRepository
}

func NewSupplierService(repo repositories.SupplierRepository) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) RegisterSupplier(name, contactInfo string) error {
	supplier := &models.Supplier{
		Nombre:       name,
		ContactoInfo: contactInfo,
	}
	return s.repo.CreateSupplier(supplier)
}

func (s *supplierService) UpdateSupplier(id uint, name, contactInfo string, rating float64) error {
	supplier, err := s.repo.GetSupplierByID(id)
	if err != nil {
		return errors.New("supplier not found")
	}

	supplier.Nombre = name
	supplier.ContactoInfo = contactInfo
	supplier.Calificacion = rating
	return s.repo.UpdateSupplier(supplier)
}

func (s *supplierService) DeleteSupplier(id uint) error {
	return s.repo.DeleteSupplier(id)
}

func (s *supplierService) GetSupplierByID(id uint) (*models.Supplier, error) {
	return s.repo.GetSupplierByID(id)
}

func (s *supplierService) GetAllSuppliers() ([]models.Supplier, error) {
	return s.repo.GetAllSuppliers()
}

func (s *supplierService) AddSupplierHistory(supplierID uint, orderID string, amount float64, description string) error {
	history := &models.SupplierHistory{
		ProveedorID: supplierID,
		OrdenID:     orderID,
		Cantidad:    amount,
		Fecha:       time.Now(),
		Descripcion: description,
	}
	return s.repo.AddSupplierHistory(history)
}

func (s *supplierService) GetSupplierHistory(supplierID uint) ([]models.SupplierHistory, error) {
	return s.repo.GetSupplierHistory(supplierID)
}
