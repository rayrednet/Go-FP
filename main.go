package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-final-project/config"
	"go-final-project/controllers"
	"go-final-project/models"
	"go-final-project/seeders"
)

func main() {

	db := config.GetDB()

	err := db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Review{})
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
	fmt.Println("Database migrated successfully!")

	seeders.SeedDatabase(db)

	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Add route for favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/images/favicon.ico")
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/categories", controllers.GetCategories)
	r.GET("/categories/create", controllers.NewCategory)
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories/:id/edit", controllers.EditCategory)
	r.POST("/categories/:id/update", controllers.UpdateCategory)
	r.POST("/categories/:id/delete", controllers.DeleteCategory)

	r.GET("/products", controllers.GetProducts)
	r.GET("/products/create", controllers.NewProduct)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id/edit", controllers.EditProduct)
	r.POST("/products/:id/update", controllers.UpdateProduct)
	r.POST("/products/:id/delete", controllers.DeleteProduct)
	r.GET("/products/:id", controllers.ShowProduct)

	r.GET("/reviews", controllers.GetReviews)
	r.GET("/reviews/create", controllers.NewReview)
	r.POST("/reviews", controllers.CreateReview)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	r.Run(":8080")
}
