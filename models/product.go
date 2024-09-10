package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string
    Stock    int
    Price    float64
    Status   string
}
