package models

import "gorm.io/gorm"

type SaleDetail struct {
    gorm.Model
    SaleID   uint
    ProductID uint
    Quantity int
    Price    float64
    Total    float64
}

type CreateSaleDetailInput struct {
    ProductID uint `json:"product_id" binding:"required"`
    Quantity  int  `json:"quantity" binding:"required,min=1"`
    Price     int  `json:"price"`
}
