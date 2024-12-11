package models

import "time"

type Supplier struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Nombre       string    `json:"name" gorm:"not null"`
	ContactoInfo string    `json:"contact_info"`
	Calificacion float64   `json:"rating"`
	Crear        time.Time `json:"created_at"`
	Actualizar   time.Time `json:"updated_at"`
}

type SupplierHistory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProveedorID uint      `json:"supplier_id" gorm:"not null"`
	OrdenID     string    `json:"order_id"`
	Cantidad    float64   `json:"amount"`
	Fecha       time.Time `json:"date"`
	Descripcion string    `json:"description"`
}
