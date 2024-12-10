package models

import "gorm.io/gorm"

type Barista struct {
    gorm.Model
    Name       string  `gorm:"not null" form:"name"`       // Barista's name
    Experience int     `gorm:"not null" form:"experience"` // Years of experience
    ProfilePic string  `gorm:"not null" form:"profile_pic"` // Profile picture URL
    Rating     float64 `gorm:"not null" form:"rating"`     // Barista's average rating
}
