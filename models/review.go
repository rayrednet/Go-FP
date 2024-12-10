package models

import "gorm.io/gorm"

type Review struct {
    gorm.Model
    CustName  string  `gorm:"not null" form:"cust_name"` // Customer name
    CustEmail string  `gorm:"not null" form:"cust_email"` // Customer email
    Review    string  `gorm:"not null" form:"review"`   // Review text
    ProductID uint    `gorm:"not null"`                 // Foreign key to Product
    Product   Product `gorm:"foreignKey:ProductID"`     // Reference to Product
}
