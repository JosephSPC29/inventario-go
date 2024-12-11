package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Usuario    string    `json:"username" gorm:"unique;not null"`
	Contraseña string    `json:"-" gorm:"not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Rol        string    `json:"role" gorm:"not null"`
	Crear      time.Time `json:"created_at"`
	Actualizar time.Time `json:"updated_at"`
}
