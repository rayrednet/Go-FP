package controllers

import (
	"net/http"
	"go-final-project/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"path/filepath"
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
	sort := c.Query("sort") // Get the sort query parameter

	// Apply sorting logic
	query := bc.DB
	if sort == "newest" {
		query = query.Order("created_at DESC") // Sort by newest
	} else if sort == "oldest" {
		query = query.Order("created_at ASC") // Sort by oldest
	}

	// Fetch the baristas from the database
	if err := query.Find(&baristas).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to fetch baristas",
		})
		return
	}

	// Pass the sorted baristas and the sort parameter to the HTML template
	c.HTML(http.StatusOK, "barista_index.html", gin.H{
		"baristas": baristas,
		"sort":     sort, // Include the current sort parameter for UI state
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

// EditBarista renders the edit form for a specific barista
func (bc *BaristaController) EditBarista(c *gin.Context) {
	id := c.Param("id")
	var barista models.Barista

	// Find the barista by ID
	if err := bc.DB.First(&barista, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Barista not found",
		})
		return
	}

	// Render the edit form with the barista data
	c.HTML(http.StatusOK, "barista_edit.html", gin.H{
		"barista": barista,
	})
}

// UpdateBarista updates a specific barista's details
func (bc *BaristaController) UpdateBarista(c *gin.Context) {
	id := c.Param("id")
	var barista models.Barista

	// Find the barista by ID
	if err := bc.DB.First(&barista, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "Barista not found",
		})
		return
	}

	// Parse form data without overwriting the existing object
	var form struct {
		Name       string `form:"Name"`
		Experience int    `form:"Experience"`
		Rating     float64 `form:"Rating"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "barista_edit.html", gin.H{
			"error": "Invalid form data",
			"barista": barista,
		})
		return
	}

	// Update fields only if present in the form
	barista.Name = form.Name
	barista.Experience = form.Experience
	barista.Rating = form.Rating

	// Handle profile picture update
	file, err := c.FormFile("ProfilePic")
	if err == nil { // If a new file is uploaded
		// Save the new file
		filePath := "./static/uploads/" + file.Filename
		if saveErr := c.SaveUploadedFile(file, filePath); saveErr != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": "Failed to save profile picture",
			})
			return
		}
		// Update the ProfilePic field
		barista.ProfilePic = "/static/uploads/" + file.Filename
	}

	// Save updated barista to the database
	if err := bc.DB.Save(&barista).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to update barista",
		})
		return
	}

	// Redirect to the barista list page
	c.Redirect(http.StatusFound, "/baristas")
}

func (bc *BaristaController) NewBarista(c *gin.Context) {
    c.HTML(http.StatusOK, "barista_create.html", nil)
}


func (bc *BaristaController) CreateBarista(c *gin.Context) {
    var barista models.Barista

    // Manually bind form fields to the Barista struct
    barista.Name = c.PostForm("Name")
    experience, err := strconv.Atoi(c.PostForm("Experience"))
    if err == nil {
        barista.Experience = experience
    }
    rating, err := strconv.ParseFloat(c.PostForm("Rating"), 64)
    if err == nil {
        barista.Rating = rating
    }

    // Handle profile picture upload
    file, err := c.FormFile("ProfilePic")
    if err == nil { // If a file is uploaded
        filename := filepath.Base(file.Filename)
        savePath := filepath.Join("static", "uploads", filename)
        if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
            c.HTML(http.StatusInternalServerError, "barista_create.html", gin.H{
                "error": "Failed to save profile picture",
            })
            return
        }
        // Save a client-accessible path to the database
        barista.ProfilePic = "/static/uploads/" + filename
    }

    // Save the new barista to the database
    if err := bc.DB.Create(&barista).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "barista_create.html", gin.H{
            "error": "Failed to create barista",
        })
        return
    }

    // Redirect to the baristas list
    c.Redirect(http.StatusFound, "/baristas")
}
