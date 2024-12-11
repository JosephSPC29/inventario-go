package repositories

import (
	"inventario-go/models"

	"gorm.io/gorm"
)

type ConfigRepository interface {
	GetSystemConfig() (*models.SystemConfig, error)
	UpdateSystemConfig(config *models.SystemConfig) error
	GetInventoryConfig() (*models.InventoryConfig, error)
	UpdateInventoryConfig(config *models.InventoryConfig) error
}

type configRepository struct {
	db *gorm.DB
}

func NewConfigRepository(db *gorm.DB) ConfigRepository {
	return &configRepository{db}
}

func (r *configRepository) GetSystemConfig() (*models.SystemConfig, error) {
	var config models.SystemConfig
	err := r.db.First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *configRepository) UpdateSystemConfig(config *models.SystemConfig) error {
	return r.db.Save(config).Error
}

func (r *configRepository) GetInventoryConfig() (*models.InventoryConfig, error) {
	var config models.InventoryConfig
	err := r.db.First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *configRepository) UpdateInventoryConfig(config *models.InventoryConfig) error {
	return r.db.Save(config).Error
}
