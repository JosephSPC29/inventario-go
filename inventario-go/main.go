package main

import (
	"inventario-go/handlers"
	"inventario-go/repositories"
	"inventario-go/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	saleRepo := repositories.NewSalesRepository(db)
	configRepo := repositories.NewConfigRepository(db)
	inventoryRepo := repositories.NewInventoryRepository(db)
	supplierRepo := repositories.NewSupplierRepository(db)
	reportRepo := repositories.NewReportRepository(db)

	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	saleService := services.NewSaleService(saleRepo)
	configService := services.NewConfigService(configRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)
	supplierService := services.NewSupplierService(supplierRepo)
	reportService := services.NewReportService(reportRepo)

	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)
	saleHandler := handlers.NewSalesHandler(saleService)
	configHandler := handlers.NewConfigHandler(configService)
	inventoryHandler := handlers.NewInventoryHandler(inventoryService)
	supplierHandler := handlers.NewSupplierHandler(supplierService)
	reportHandler := handlers.NewReportsHandler(reportService)

	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/entry", userHandler.Register)
		userGroup.POST("/welcome", userHandler.Login)
	}

	productGroup := r.Group("/products")
	{
		productGroup.GET("/", productHandler.GetProducts)
		productGroup.POST("/", productHandler.RegisterProduct)
		productGroup.PUT("/:id", productHandler.UpdateProduct)
		productGroup.DELETE("/:id", productHandler.DeleteProduct)
	}

	saleGroup := r.Group("/sales")
	{
		saleGroup.GET("/entry", saleHandler.RegisterSale)
		saleGroup.GET("/:id", saleHandler.GetTopSellingProducts)
	}

	configGroup := r.Group("/config")
	{
		configGroup.GET("/system", configHandler.GetSystemConfig)
		configGroup.PUT("/system", configHandler.UpdateSystemConfig)
		configGroup.GET("/inventory", configHandler.GetInventoryConfig)
		configGroup.PUT("/inventory", configHandler.UpdateInventoryConfig)
	}

	inventoryGroup := r.Group("/inventory")
	{
		inventoryGroup.GET("/state", inventoryHandler.GetLowStockAlerts)
		inventoryGroup.POST("/entry", inventoryHandler.RegisterStock)
		inventoryGroup.PUT("/exit", inventoryHandler.UpdateStock)
		inventoryGroup.GET("/movements", inventoryHandler.GetExpiredProducts)
	}

	supplierGroup := r.Group("/suppliers")
	{
		supplierGroup.GET("/", supplierHandler.GetAllSuppliers)
		supplierGroup.POST("/", supplierHandler.RegisterSupplier)
		supplierGroup.PUT("/:id", supplierHandler.UpdateSupplier)
	}

	reportGroup := r.Group("/reports")
	{
		reportGroup.GET("/inventory", reportHandler.GetInventoryReport)
		reportGroup.GET("/sales", reportHandler.GetSalesReport)
		reportGroup.GET("/movements", reportHandler.GetInventoryMovements)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Servidor en funcionamiento"})
	})

}
