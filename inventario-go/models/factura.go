package models

import "time"

type Invoice struct {
	ID      uint      `json:"id" gorm:"primaryKey"`
	VentaID uint      `json:"sale_id" gorm:"not null"`
	Fecha   time.Time `json:"invoice_date"`
	Total   float64   `json:"total"`
}
