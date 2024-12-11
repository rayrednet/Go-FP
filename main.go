package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"go-final-project/config"
	"go-final-project/models"
	"go-final-project/routes"
	"go-final-project/seeders"
)

func main() {
	db := config.GetDB()

	// AutoMigrate database
	err := db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Review{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated successfully!")

	// Seed the database
	seeders.SeedDatabase(db)

	// Initialize Gin router
	r := gin.Default()

	// Register the `hasPrefix` function
	r.SetFuncMap(map[string]interface{}{
		"hasPrefix": strings.HasPrefix,
	})

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Register routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":8080")
}
