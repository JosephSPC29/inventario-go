package handlers

import (
	"inventario-go/models"
	"inventario-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SalesHandler struct {
	service services.SalesService
}

func NewSalesHandler(service services.SalesService) *SalesHandler {
	return &SalesHandler{service}
}

func (h *SalesHandler) RegisterSale(c *gin.Context) {
	var req struct {
		Cliente    string            `json:"customer" binding:"required"`
		MetodoPago string            `json:"payment_method" binding:"required"`
		Items      []models.SaleItem `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saleID, err := h.service.RegisterSale(req.Cliente, req.MetodoPago, req.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Sale registered successfully", "sale_id": saleID})
}

func (h *SalesHandler) GetTopSellingProducts(c *gin.Context) {
	topProducts, err := h.service.GetTopSellingProducts(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topProducts)
}
