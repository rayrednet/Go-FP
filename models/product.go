package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name          string
    Description   string
    Price         float64
    CategoryID    uint
    Category      Category
    AvailableHour string
    Rating        float64
    Stock         int
    ImagePath     string  `gorm:"type:varchar(255)"`
}