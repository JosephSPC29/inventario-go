package repositories

import (
	"inventario-go/models"
	"time"

	"gorm.io/gorm"
)

type ReportsRepository interface {
	GetInventoryState() ([]models.Product, error)
	GetProductMovements(startDate, endDate time.Time) ([]models.InventoryMovement, error)
	GetSalesSummary(startDate, endDate time.Time) ([]SalesSummary, error)
}

type reportsRepository struct {
	db *gorm.DB
}

func NewReportsRepository(db *gorm.DB) ReportsRepository {
	return &reportsRepository{db}
}

func (r *reportsRepository) GetInventoryState() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *reportsRepository) GetProductMovements(startDate, endDate time.Time) ([]models.InventoryMovement, error) {
	var movements []models.InventoryMovement
	err := r.db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&movements).Error
	return movements, err
}

type SalesSummary struct {
	Fecha       time.Time `json:"date"`
	TotalVentas float64   `json:"total_sales"`
	TotalItems  int       `json:"total_items"`
}

func (r *reportsRepository) GetSalesSummary(startDate, endDate time.Time) ([]SalesSummary, error) {
	var summary []SalesSummary
	err := r.db.Raw(`
		SELECT 
			DATE(created_at) AS date, 
			SUM(total_amount) AS total_sales, 
			SUM(quantity) AS total_items
		FROM sales
		JOIN sale_items ON sales.id = sale_items.sale_id
		WHERE created_at BETWEEN ? AND ?
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) ASC
	`, startDate, endDate).Scan(&summary).Error
	return summary, err
}
