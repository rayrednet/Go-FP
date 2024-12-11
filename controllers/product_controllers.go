package controllers

import (
	"go-final-project/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "product_index.html", gin.H{
		"products": products,
	})
}

func NewProduct(c *gin.Context) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "product_create.html", gin.H{
		"categories": categories,
		"product":    nil,
	})
}

func CreateProduct(c *gin.Context) {
	var input struct {
		Name          string  `form:"name" binding:"required"`
		Description   string  `form:"description" binding:"required"`
		Price         float64 `form:"price" binding:"required"`
		CategoryID    uint    `form:"category_id" binding:"required"`
		AvailableHour string  `form:"available_hour" binding:"required"`
		Rating        float64 `form:"rating" binding:"required"`
		Stock         int     `form:"stock" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:          input.Name,
		Description:   input.Description,
		Price:         input.Price,
		CategoryID:    input.CategoryID,
		AvailableHour: input.AvailableHour,
		Rating:        input.Rating,
		Stock:         input.Stock,
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/products")
}

func EditProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Normalize image path
	if product.ImagePath != "" {
		product.ImagePath = filepath.ToSlash(product.ImagePath)
	}

	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "product_edit.html", gin.H{
		"product":    product,
		"categories": categories,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input struct {
		Name          string  `form:"name" binding:"required"`
		Description   string  `form:"description" binding:"required"`
		Price         float64 `form:"price" binding:"required"`
		CategoryID    uint    `form:"category_id" binding:"required"`
		AvailableHour string  `form:"available_hour" binding:"required"`
		Rating        float64 `form:"rating" binding:"required"`
		Stock         int     `form:"stock" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update product fields
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.CategoryID = input.CategoryID
	product.AvailableHour = input.AvailableHour
	product.Rating = input.Rating
	product.Stock = input.Stock

	// Handle image upload
	file, err := c.FormFile("image")
	if err == nil { // File is uploaded
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join("static", "uploads", "products", filename)

		// Save the file to the uploads folder
		if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
			return
		}

		// Update the image path in the database
		product.ImagePath = filepath.ToSlash(savePath)
	}

	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/products")
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/products")
}

func ShowProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := db.Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "product_detail.html", gin.H{
		"product": product,
	})
}
