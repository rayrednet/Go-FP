package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-final-project/config"
	"go-final-project/models"

)

func main() {
	db := config.GetDB()

	err := db.AutoMigrate(&models.Category{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated successfully!")

	r := gin.Default()


	r.LoadHTMLGlob("templates/*")


	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	r.Run(":8080")
}
