package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name          string  `gorm:"not null" form:"name"`
    Description   string  `gorm:"not null" form:"description"`
    Price         float64 `gorm:"not null" form:"price"`
    CategoryID    uint    
    Category      Category `gorm:"foreignKey:CategoryID"` 
    AvailableHour string  `gorm:"not null" form:"available_hour"` 
    Rating        float64 `gorm:"not null" form:"rating"`       
    Stock         int     `gorm:"not null" form:"stock"`        
}
