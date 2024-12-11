package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Nombre      string    `json:"name" gorm:"not null"`
	Descripcion string    `json:"description"`
	Categoria   string    `json:"category" gorm:"not null"`
	Precio      float64   `json:"price" gorm:"not null"`
	Crear       time.Time `json:"created_at"`
	Actualizar  time.Time `json:"updated_at"`
}
