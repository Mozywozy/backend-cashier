package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "cashier-app/models"
    "cashier-app/config"
)

func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := config.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving products"})
        return
    }
    c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    
    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }

    var input models.Product
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    product.Name = input.Name
    product.Stock = input.Stock
    product.Price = input.Price
    product.Status = input.Status

    if err := config.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating product"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
