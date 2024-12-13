package routes

import (
	"github.com/gin-gonic/gin"
	"go-final-project/config"
	"go-final-project/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	// Home route
	r.GET("/", controllers.HomeHandler)

	// Category routes
	r.GET("/categories", controllers.GetCategories)
	r.GET("/categories/create", controllers.NewCategory)
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories/:id/edit", controllers.EditCategory)
	r.POST("/categories/:id/update", controllers.UpdateCategory)
	r.POST("/categories/:id/delete", controllers.DeleteCategory)

	// Product routes
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/create", controllers.NewProduct)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id/edit", controllers.EditProduct)
	r.POST("/products/:id/update", controllers.UpdateProduct)
	r.POST("/products/:id/delete", controllers.DeleteProduct)
	r.GET("/products/:id", controllers.ShowProduct)

	// Barista routes
	db := config.GetDB()
	baristaController := controllers.NewBaristaController(db)

	r.GET("/baristas", baristaController.GetAllBaristas)         // Get all baristas with optional sorting
	r.POST("/baristas/:id/delete", baristaController.DeleteBarista) // Delete a barista
	r.GET("/baristas/:id/edit", baristaController.EditBarista)      // Show edit form
	r.POST("/baristas/:id/update", baristaController.UpdateBarista) // Handle update request
	r.GET("/baristas/create", baristaController.NewBarista)         // Show create form
	r.POST("/baristas", baristaController.CreateBarista)            // Handle form submission


	// Review routes
	r.GET("/reviews", controllers.GetReviews)
	r.GET("/reviews/create", controllers.NewReview)
	r.POST("/reviews", controllers.CreateReview)
	r.POST("/reviews/:id/delete", controllers.DeleteReview)
	r.GET("/reviews/:id/edit", controllers.EditReview)
	r.POST("/reviews/:id/update", controllers.UpdateReview)
}
