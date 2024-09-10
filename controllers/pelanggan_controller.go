package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "cashier-app/models"
    "cashier-app/config"
)

func GetTransactions(c *gin.Context) {
    userID, _ := c.Get("user_id")
    var sales []models.Sale
    if err := config.DB.Where("user_id = ?", userID).Preload("Details").Find(&sales).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving transactions"})
        return
    }
    c.JSON(http.StatusOK, sales)
}
