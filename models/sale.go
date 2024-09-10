package models

import "gorm.io/gorm"

type Sale struct {
    gorm.Model
    UserID  uint
    Total   float64
    Details   []SaleDetail   `gorm:"foreignKey:SaleID"`
}

type CreateSaleInput struct {
    UserID   uint                   `json:"user_id" binding:"required"`
    Products []CreateSaleDetailInput `json:"products" binding:"required"`
}
