package models

import "time"

type Inventory struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	ProductoID      uint       `json:"product_id" gorm:"not null"`
	Stock           int        `json:"stock" gorm:"not null"`
	StockMinimo     int        `json:"minimum_stock" gorm:"not null"`
	FechaExpiracion *time.Time `json:"expiration_date"`
	Crear           time.Time  `json:"created_at"`
	Actualizar      time.Time  `json:"updated_at"`
}
