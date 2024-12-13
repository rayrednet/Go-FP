package models

import "gorm.io/gorm"

type Barista struct {
    gorm.Model
    Name       string    `gorm:"not null" form:"name"`
    Experience int       `gorm:"not null" form:"experience"`
    ProfilePic string    `gorm:"not null" form:"profile_pic"`
    Rating     float64   `gorm:"not null" form:"rating"`
    Specialty  string    `gorm:"not null" form:"specialty"`
    Products   []Product `gorm:"foreignKey:BaristaID"`
}