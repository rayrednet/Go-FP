package controllers

import (
	"net/http"
	"go-final-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BaristaController structure
type BaristaController struct {
	DB *gorm.DB
}

// NewBaristaController initializes a new BaristaController
func NewBaristaController(db *gorm.DB) *BaristaController {
	return &BaristaController{DB: db}
}

// GetAllBaristas fetches all baristas and displays them in an HTML template
func (bc *BaristaController) GetAllBaristas(c *gin.Context) {
	var baristas []models.Barista

	// Fetch all baristas from the database
	if err := bc.DB.Find(&baristas).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to fetch baristas",
		})
		return
	}

	// Pass the baristas to the HTML template
	c.HTML(http.StatusOK, "barista_index.html", gin.H{
		"baristas": baristas,
	})
}

// DeleteBarista deletes a specific barista by ID
func (bc *BaristaController) DeleteBarista(c *gin.Context) {
	id := c.Param("id") // Get the barista ID from the URL

	// Find and delete the barista
	if err := bc.DB.Delete(&models.Barista{}, id).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to delete barista",
		})
		return
	}

	// Redirect back to the baristas list
	c.Redirect(http.StatusFound, "/baristas")
}
