package models

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name        string `gorm:"not null" form:"name"`        // Category name
    Description string `gorm:"not null" form:"description"` // Category description
    ImagePath   string `gorm:"not null" form:"image_path"`  // Path to category image
}
