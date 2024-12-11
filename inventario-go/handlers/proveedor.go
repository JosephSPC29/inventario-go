package handlers

import (
	"inventario-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	service services.SupplierService
}

func NewSupplierHandler(service services.SupplierService) *SupplierHandler {
	return &SupplierHandler{service}
}

func (h *SupplierHandler) RegisterSupplier(c *gin.Context) {
	var req struct {
		Nombre       string `json:"name" binding:"required"`
		ContactoInfo string `json:"contact_info"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterSupplier(req.Nombre, req.ContactoInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Supplier registered successfully"})
}

func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	var req struct {
		Nombre       string  `json:"name" binding:"required"`
		ContactoInfo string  `json:"contact_info"`
		Calificacion float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid supplier ID"})
		return
	}

	err = h.service.UpdateSupplier(uint(id), req.Nombre, req.ContactoInfo, req.Calificacion)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated successfully"})
}

func (h *SupplierHandler) GetAllSuppliers(c *gin.Context) {
	suppliers, err := h.service.GetAllSuppliers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}
