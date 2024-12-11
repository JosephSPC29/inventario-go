package services

import (
	"inventario-go/models"
	"inventario-go/repositories"
)

type ConfigService interface {
	GetSystemConfig() (*models.SystemConfig, error)
	UpdateSystemConfig(config *models.SystemConfig) error
	GetInventoryConfig() (*models.InventoryConfig, error)
	UpdateInventoryConfig(config *models.InventoryConfig) error
}

type configService struct {
	repo repositories.ConfigRepository
}

func NewConfigService(repo repositories.ConfigRepository) ConfigService {
	return &configService{repo}
}

func (s *configService) GetSystemConfig() (*models.SystemConfig, error) {
	return s.repo.GetSystemConfig()
}

func (s *configService) UpdateSystemConfig(config *models.SystemConfig) error {
	return s.repo.UpdateSystemConfig(config)
}

func (s *configService) GetInventoryConfig() (*models.InventoryConfig, error) {
	return s.repo.GetInventoryConfig()
}

func (s *configService) UpdateInventoryConfig(config *models.InventoryConfig) error {
	return s.repo.UpdateInventoryConfig(config)
}
