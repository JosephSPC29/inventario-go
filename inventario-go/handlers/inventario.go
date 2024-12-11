package handlers

import (
	"inventario-go/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	service services.InventoryService
}

func NewInventoryHandler(service services.InventoryService) *InventoryHandler {
	return &InventoryHandler{service}
}

func (h *InventoryHandler) RegisterStock(c *gin.Context) {
	var req struct {
		ProductoID      uint       `json:"product_id" binding:"required"`
		Stock           int        `json:"stock" binding:"required"`
		StockMinimo     int        `json:"minimum_stock" binding:"required"`
		FechaExpiracion *time.Time `json:"expiration_date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterStock(req.ProductoID, req.Stock, req.StockMinimo, req.FechaExpiracion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Stock registered successfully"})
}

func (h *InventoryHandler) UpdateStock(c *gin.Context) {
	var req struct {
		ProductoID  uint `json:"product_id" binding:"required"`
		StockCambio int  `json:"stock_change" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateStock(req.ProductoID, req.StockCambio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
}

func (h *InventoryHandler) GetLowStockAlerts(c *gin.Context) {
	alerts, err := h.service.GetLowStockAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alerts)
}

func (h *InventoryHandler) GetExpiredProducts(c *gin.Context) {
	expiredProducts, err := h.service.GetExpiredProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expiredProducts)
}
