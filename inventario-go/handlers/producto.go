package handlers

import (
	"inventario-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) RegisterProduct(c *gin.Context) {
	var req struct {
		Nombre      string  `json:"name" binding:"required"`
		Descripcion string  `json:"description"`
		Categoria   string  `json:"category" binding:"required"`
		Precio      float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RegisterProduct(req.Nombre, req.Descripcion, req.Categoria, req.Precio)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product registered successfully"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Nombre      string  `json:"name" binding:"required"`
		Descripcion string  `json:"description"`
		Categoria   string  `json:"category" binding:"required"`
		Precio      float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateProduct(uint(id), req.Nombre, req.Descripcion, req.Categoria, req.Precio)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	filter := map[string]interface{}{
		"name":      c.Query("name"),
		"category":  c.Query("category"),
		"min_price": c.Query("min_price"),
		"max_price": c.Query("max_price"),
	}

	products, err := h.service.GetProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
