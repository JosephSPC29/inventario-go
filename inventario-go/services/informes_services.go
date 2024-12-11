package services

import (
	"inventario-go/models"
	"inventario-go/repositories"
	"time"
)

type ReportsService interface {
	GetInventoryReport() ([]models.Product, error)
	GetInventoryMovements(startDate, endDate time.Time) ([]models.InventoryMovement, error)
	GetSalesReport(startDate, endDate time.Time) ([]repositories.SalesSummary, error)
}

type reportsService struct {
	repo repositories.ReportsRepository
}

func NewReportService(repo repositories.ReportsRepository) ReportsService {
	return &reportsService{repo}
}

func (s *reportsService) GetInventoryReport() ([]models.Product, error) {
	return s.repo.GetInventoryState()
}

func (s *reportsService) GetInventoryMovements(startDate, endDate time.Time) ([]models.InventoryMovement, error) {
	return s.repo.GetProductMovements(startDate, endDate)
}

func (s *reportsService) GetSalesReport(startDate, endDate time.Time) ([]repositories.SalesSummary, error) {
	return s.repo.GetSalesSummary(startDate, endDate)
}
