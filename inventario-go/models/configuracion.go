package models

import "time"

type SystemConfig struct {
	ID                          uint      `json:"id" gorm:"primaryKey"`
	Inventario                  int       `json:"inventory_threshold"`
	NotificacionesHabilitadas   bool      `json:"notification_enabled"`
	NotificacionEmail           string    `json:"notification_email"`
	NotificationFrecuencia      string    `json:"notification_frequency"`
	PredeterminacionConcurrente string    `json:"default_currency"`
	Crear                       time.Time `json:"created_at"`
	Actualizar                  time.Time `json:"updated_at"`
}

type InventoryConfig struct {
	ID                                  uint      `json:"id" gorm:"primaryKey"`
	ReordenarProductos                  int       `json:"reorder_level"`
	Actualizaciones                     bool      `json:"automatic_stock_updates"`
	AdvertenciaHabilitacionNotificacion bool      `json:"stock_warning_enabled"`
	AdvertenciaStockBajo                int       `json:"stock_warning_threshold"`
	Creae                               time.Time `json:"created_at"`
	Actualizar                          time.Time `json:"updated_at"`
}
