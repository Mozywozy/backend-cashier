package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "cashier-app/models"
    "cashier-app/config"
)

func GetSales(c *gin.Context) {
    var sales []models.Sale
    if err := config.DB.Preload("Details").Find(&sales).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving sales"})
        return
    }
    c.JSON(http.StatusOK, sales)
}

func GetSale(c *gin.Context) {
    id := c.Param("id")
    var sale models.Sale
    if err := config.DB.Preload("Details").First(&sale, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Sale not found"})
        return
    }
    c.JSON(http.StatusOK, sale)
}

func CreateSale(c *gin.Context) {
    var input models.CreateSaleInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hitung total harga
    totalPrice := 0
    for _, product := range input.Products {
        var prod models.Product
        if err := config.DB.First(&prod, product.ProductID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }
        if prod.Stock < product.Quantity {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product: " + prod.Name})
            return
        }

        totalPrice += int(prod.Price) * product.Quantity
        prod.Stock -= product.Quantity
        config.DB.Save(&prod) 
    }

    // Simpan transaksi
    sale := models.Sale{
        UserID:   input.UserID,
        Total:    float64(totalPrice),
    }

    if err := config.DB.Create(&sale).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale"})
        return
    }

    // Simpan detail transaksi
    for _, product := range input.Products {
        saleDetail := models.SaleDetail{
            SaleID:    sale.ID,
            ProductID: product.ProductID,
            Quantity:  product.Quantity,
            Price:     float64(product.Price),
            Total:     float64(product.Price) * float64(product.Quantity),
        }
        config.DB.Create(&saleDetail)
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func UpdateSale(c *gin.Context) {
    id := c.Param("id")
    var sale models.Sale
    if err := config.DB.First(&sale, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Sale not found"})
        return
    }

    var input struct {
        UserID   uint                `json:"user_id"`
        Products []models.SaleDetail `json:"products"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    total := 0.0
    for _, product := range input.Products {
        prod := models.Product{}
        if err := config.DB.First(&prod, product.ProductID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
            return
        }
        if prod.Stock < product.Quantity {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient stock for product"})
            return
        }
        total += prod.Price * float64(product.Quantity)
    }

    sale.UserID = input.UserID
    sale.Total = total

    if err := config.DB.Save(&sale).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating sale"})
        return
    }

    config.DB.Delete(&models.SaleDetail{}, "sale_id = ?", id)
    for _, product := range input.Products {
        config.DB.Create(&models.SaleDetail{
            SaleID:    sale.ID,
            ProductID: product.ProductID,
            Quantity:  product.Quantity,
            Price:     product.Price,
            Total:     product.Total,
        })
    }

    c.JSON(http.StatusOK, sale)
}

func DeleteSale(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.Sale{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Sale not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Sale deleted"})
}
