package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB == nil {
		
		dsn := "root:helsasp@tcp(127.0.0.1:3306)/fp_pbkk?charset=utf8mb4&parseTime=True&loc=Local"
		
		
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}


		fmt.Println("Connected to the database successfully")
	}
	return DB
}
