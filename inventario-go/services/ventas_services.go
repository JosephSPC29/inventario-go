package services

import (
	"inventario-go/models"
	"inventario-go/repositories"
)

type SalesService interface {
	RegisterSale(customer string, paymentMethod string, items []models.SaleItem) (uint, error)
	GetSaleDetails(id uint) (*models.Sale, error)
	GetAllSales() ([]models.Sale, error)
	GetTopSellingProducts(limit int) ([]models.SaleItem, error)
}

type salesService struct {
	repo repositories.SalesRepository
}

func NewSaleService(repo repositories.SalesRepository) SalesService {
	return &salesService{repo}
}

func (s *salesService) RegisterSale(customer string, paymentMethod string, items []models.SaleItem) (uint, error) {
	sale := &models.Sale{
		Cliente:       customer,
		MetodoPago:    paymentMethod,
		TotalCantidad: calculateTotal(items),
		Items:         items,
	}

	if err := s.repo.CreateSale(sale); err != nil {
		return 0, err
	}
	return sale.ID, nil
}

func (s *salesService) GetSaleDetails(id uint) (*models.Sale, error) {
	return s.repo.GetSaleByID(id)
}

func (s *salesService) GetAllSales() ([]models.Sale, error) {
	return s.repo.GetAllSales()
}

func (s *salesService) GetTopSellingProducts(limit int) ([]models.SaleItem, error) {
	return s.repo.GetTopSellingProducts(limit)
}

func calculateTotal(items []models.SaleItem) float64 {
	total := 0.0
	for _, item := range items {
		total += float64(item.Cantidad) * item.Precio
	}
	return total
}
