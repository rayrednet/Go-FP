package models

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name        string `gorm:"not null" form:"name"`        
    Description string `gorm:"not null" form:"description"` 
    ImagePath   string `gorm:"not null" form:"image_path"`  
}
