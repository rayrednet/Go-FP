package controllers

import (
	"go-final-project/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Preload("Category").Preload("Barista").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "product_index.html", gin.H{
		"products": products,
	})
}

func NewProduct(c *gin.Context) {
	var categories []models.Category
	var baristas []models.Barista

	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Find(&baristas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "product_create.html", gin.H{
		"categories": categories,
		"baristas":   baristas,
		"product":    nil,
	})
}

func CreateProduct(c *gin.Context) {
	var input struct {
		Name          string  `form:"name" binding:"required"`
		Description   string  `form:"description" binding:"required"`
		Price         float64 `form:"price" binding:"required"`
		CategoryID    uint    `form:"category_id" binding:"required"`
		BaristaID     uint    `form:"barista_id" binding:"required"`
		AvailableHour string  `form:"available_hour" binding:"required"`
		Rating        float64 `form:"rating" binding:"required"`
		Stock         int     `form:"stock" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	filename := filepath.Base(file.Filename)
	savePath := filepath.Join("static", "uploads", "products", filename)

	if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	// Create product instance
	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		CategoryID:    input.CategoryID,
		BaristaID:     input.BaristaID,
		AvailableHour: input.AvailableHour,
		Rating:        input.Rating,
		Stock:         input.Stock,
		ImagePath:     filepath.ToSlash(savePath), // Normalize the path
	}

	// Save to database
	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/products")
}

func EditProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Preload("Category").Preload("Barista").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Normalize image path
	if product.ImagePath != "" {
		product.ImagePath = filepath.ToSlash(product.ImagePath)
	}

	var categories []models.Category
	var baristas []models.Barista

	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Find(&baristas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "product_edit.html", gin.H{
		"product":    product,
		"categories": categories,
		"baristas":   baristas,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Fetch the existing product from the database
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Parse form data
	var input struct {
		Name          string  `form:"name"`
		Description   string  `form:"description"`
		Price         float64 `form:"price"`
		CategoryID    uint    `form:"category_id"`
		AvailableHour string  `form:"available_hour"`
		Rating        float64 `form:"rating"`
		Stock         int     `form:"stock"`
		BaristaID     uint    `form:"barista_id"`
	}

	// Only bind and validate fields included in the form
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided in the form
	if input.Name != "" {
		product.Name = input.Name
	}
	if input.Description != "" {
		product.Description = input.Description
	}
	if input.Price != 0 {
		product.Price = input.Price
	}
	if input.CategoryID != 0 {
		product.CategoryID = input.CategoryID
	}
	if input.AvailableHour != "" {
		product.AvailableHour = input.AvailableHour
	}
	if input.Rating != 0 {
		product.Rating = input.Rating
	}
	if input.Stock != 0 {
		product.Stock = input.Stock
	}
	if input.BaristaID != 0 {
		product.BaristaID = input.BaristaID
	}

	// Handle image upload
	file, err := c.FormFile("image")
	if err == nil { // If a file is uploaded
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join("static", "uploads", "products", filename)

		// Save the uploaded file
		if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}

		// Update the image path
		product.ImagePath = filepath.ToSlash(savePath)
	}

	// Save the updated product to the database
	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/products")
}

func ShowProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Preload("Category").Preload("Barista").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "product_detail.html", gin.H{
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/products")
}
