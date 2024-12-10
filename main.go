package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-final-project/config"
	"go-final-project/models"
	"go-final-project/seeders"
	"go-final-project/routes"
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

	// Serve static files
	r.Static("/static", "./static") // Ensure this is only defined here

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Register routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":8080")
}