package controllers

import (
	"go-final-project/models"
	"go-final-project/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReviews(c *gin.Context) {
	var reviews []models.Review
	db := config.GetDB()

	// Get the sort query parameter
	sort := c.Query("sort")

	// Set the default order
	order := "created_at desc"
	if sort == "oldest" {
		order = "created_at asc"
	}

	// Fetch reviews with sorting
	if err := db.Preload("Product").Order(order).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Pass the reviews and current sort option to the template
	c.HTML(http.StatusOK, "review_index.html", gin.H{
		"reviews": reviews,
		"sort":    sort,
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

// Fetch the review data and render the edit form
func EditReview(c *gin.Context) {
    id := c.Param("id")
    var review models.Review
    db := config.GetDB()

    // Find the review by ID
    if err := db.Preload("Product").First(&review, id).Error; err != nil {
        c.HTML(http.StatusNotFound, "error.html", gin.H{
            "error": "Review not found",
        })
        return
    }

    var products []models.Product
    if err := db.Find(&products).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "error.html", gin.H{
            "error": "Failed to fetch products",
        })
        return
    }

    // Render the edit form
    c.HTML(http.StatusOK, "review_edit.html", gin.H{
        "review":   review,
        "products": products,
    })
}

// Handle review updates
func UpdateReview(c *gin.Context) {
    id := c.Param("id")
    var review models.Review
    db := config.GetDB()

    // Find the review by ID
    if err := db.First(&review, id).Error; err != nil {
        c.HTML(http.StatusNotFound, "error.html", gin.H{
            "error": "Review not found",
        })
        return
    }

    // Parse updated data from form
    var input struct {
        CustName  string `form:"cust_name" binding:"required"`
        CustEmail string `form:"cust_email" binding:"required,email"`
        Review    string `form:"review" binding:"required"`
        ProductID uint   `form:"product_id" binding:"required"`
    }

    if err := c.ShouldBind(&input); err != nil {
        c.HTML(http.StatusBadRequest, "review_edit.html", gin.H{
            "error":   "Invalid form data",
            "review":  review,
            "products": nil,
        })
        return
    }

    // Update the review
    review.CustName = input.CustName
    review.CustEmail = input.CustEmail
    review.Review = input.Review
    review.ProductID = input.ProductID

    if err := db.Save(&review).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "error.html", gin.H{
            "error": "Failed to update review",
        })
        return
    }

    // Redirect to reviews page
    c.Redirect(http.StatusFound, "/reviews")
}
