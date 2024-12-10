package models

import "gorm.io/gorm"

type Review struct {
    gorm.Model
    CustName  string  `gorm:"not null" form:"cust_name"` 
    CustEmail string  `gorm:"not null" form:"cust_email"` 
    Review    string  `gorm:"not null" form:"review"`   
    ProductID uint    `gorm:"not null"`                 
    Product   Product `gorm:"foreignKey:ProductID"`     
}
