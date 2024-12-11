package models

import "time"

type Sale struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Cliente       string     `json:"customer" gorm:"not null"`
	TotalCantidad float64    `json:"total_amount" gorm:"not null"`
	MetodoPago    string     `json:"payment_method"`
	Crear         time.Time  `json:"created_at"`
	Actualizar    time.Time  `json:"updated_at"`
	Items         []SaleItem `json:"items" gorm:"foreignKey:SaleID"`
}

type SaleItem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	VentaID    uint    `json:"sale_id" gorm:"not null"`
	ProductoID uint    `json:"product_id" gorm:"not null"`
	Cantidad   int     `json:"quantity" gorm:"not null"`
	Precio     float64 `json:"price" gorm:"not null"`
}
