package controllers

import (
	"go-final-project/models"
	"go-final-project/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Display all reviews
func GetReviews(c *gin.Context) {
	var reviews []models.Review
	db := config.GetDB()

	// Preload Product to include related product data
	if err := db.Preload("Product").Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "review_index.html", gin.H{
		"reviews": reviews,
	})
}

// Show form to create a new review
func NewReview(c *gin.Context) {
	var products []models.Product
	db := config.GetDB()

	// Fetch all products for the product dropdown
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "review_create.html", gin.H{
		"products": products,
	})
}

// Handle review creation
func CreateReview(c *gin.Context) {
	var input struct {
		CustName  string `form:"cust_name" binding:"required"`
		CustEmail string `form:"cust_email" binding:"required,email"`
		Review    string `form:"review" binding:"required"`
		ProductID uint   `form:"product_id" binding:"required"`
	}

	// Bind input from the form
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Review{
		CustName:  input.CustName,
		CustEmail: input.CustEmail,
		Review:    input.Review,
		ProductID: input.ProductID,
	}

	db := config.GetDB()

	// Save the review to the database
	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to the review list
	c.Redirect(http.StatusFound, "/reviews")
}

// Delete a review by ID
func DeleteReview(c *gin.Context) {
    id := c.Param("id")
    db := config.GetDB()

    // Delete the review by ID
    if err := db.Delete(&models.Review{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
        return
    }

    // Redirect to the reviews list
    c.Redirect(http.StatusFound, "/reviews")
}
