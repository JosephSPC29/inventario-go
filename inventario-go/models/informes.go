package models

import "time"

type InventoryMovement struct {
	ID             uint      `gorm:"primaryKey"`
	ProductoID     uint      `json:"product_id"`
	Cantidad       int       `json:"quantity"`
	TypoMovimiento string    `json:"movement_type"`
	Crear          time.Time `json:"created_at"`
	Actualizar     time.Time `json:"updated_at"`
}
