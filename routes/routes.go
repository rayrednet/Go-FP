package routes

import (
	"github.com/gin-gonic/gin"
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

	// Review routes
	r.GET("/reviews", controllers.GetReviews)
	r.GET("/reviews/create", controllers.NewReview)
	r.POST("/reviews", controllers.CreateReview)
}
