package controllers

import (
	"go-final-project/config"
	"go-final-project/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var db = config.GetDB()

// GetCategories fetches all categories and renders the category list page
func GetCategories(c *gin.Context) {
	var categories []models.Category

	// Get the sort query parameter
	sort := c.Query("sort")

	// Set the default order
	order := "created_at desc"
	if sort == "oldest" {
		order = "created_at asc"
	}

	// Fetch categories with sorting
	if err := db.Order(order).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Pass the categories and current sort option to the template
	c.HTML(http.StatusOK, "category_index.html", gin.H{
		"categories": categories,
		"sort":       sort,
	})
}


// NewCategory renders the page for creating a new category
func NewCategory(c *gin.Context) {
	c.HTML(http.StatusOK, "category_create.html", nil)
}

// CreateCategory handles category creation, including image upload
func CreateCategory(c *gin.Context) {
	var category models.Category

	// Bind form data
	category.Name = c.PostForm("name")
	category.Description = c.PostForm("description")

	// Handle image upload
	file, err := c.FormFile("image")
	if err == nil { // If a file is uploaded
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join("static", "uploads", "categories", filename)

		// Save the uploaded file
		if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
			c.HTML(http.StatusInternalServerError, "category_create.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		// Set the image path
		category.ImagePath = filepath.ToSlash(savePath) // Normalize path to forward slashes
	}

	// Create the category in the database
	if err := db.Create(&category).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "category_create.html", gin.H{
			"error": "Failed to create category",
		})
		return
	}

	// Redirect to the categories list
	c.Redirect(http.StatusFound, "/categories")
}

// EditCategory renders the edit form for a specific category
func EditCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.HTML(http.StatusOK, "category_edit.html", gin.H{"category": category})
}

// UpdateCategory handles category updates, including image upload
func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	// Find the existing category by ID
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.HTML(http.StatusNotFound, "category_edit.html", gin.H{
			"error":    "Category not found",
			"category": category,
		})
		return
	}

	// Bind form data
	category.Name = c.PostForm("name")
	category.Description = c.PostForm("description")

	// Handle image upload
	file, err := c.FormFile("image")
	if err == nil { // If a file is uploaded
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join("static", "uploads", "categories", filename)

		// Save the uploaded file
		if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
			c.HTML(http.StatusInternalServerError, "category_edit.html", gin.H{
				"error":    "Failed to upload image",
				"category": category,
			})
			return
		}

		// Update the image path in the database
		category.ImagePath = filepath.ToSlash(savePath) // Normalize path to forward slashes
	}

	// Save the updated category
	if err := db.Save(&category).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "category_edit.html", gin.H{
			"error":    "Failed to update category",
			"category": category,
		})
		return
	}

	// Redirect to the categories list
	c.Redirect(http.StatusFound, "/categories")
}

// DeleteCategory deletes a specific category by ID
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.Redirect(http.StatusFound, "/categories")
}
