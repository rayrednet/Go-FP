package controllers
import (
	"go-final-project/models"
	"go-final-project/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var db = config.GetDB()


func GetCategories(c *gin.Context) {
	var categories []models.Category
	
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	

	c.HTML(http.StatusOK, "category_index.html", gin.H{"categories": categories})
}


func NewCategory(c *gin.Context) {
	c.HTML(http.StatusOK, "category_create.html", nil)
}


func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBind(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/categories")
}


func EditCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.HTML(http.StatusOK, "category_edit.html", gin.H{"category": category})
}


func UpdateCategory(c *gin.Context) {
    var category models.Category
    id := c.Param("id")

    // Find the existing category by ID
    if err := db.Where("id = ?", id).First(&category).Error; err != nil {
        c.HTML(http.StatusNotFound, "category_edit.html", gin.H{
            "error": "Category not found",
            "category": category,
        })
        return
    }

    // Bind form data
    var form struct {
        Name        string `form:"name"`
        Description string `form:"description"`
    }
    if err := c.ShouldBind(&form); err != nil {
        c.HTML(http.StatusBadRequest, "category_edit.html", gin.H{
            "error": "Invalid form data",
            "category": category,
        })
        return
    }

    // Update fields
    category.Name = form.Name
    category.Description = form.Description

    // Handle image upload
    file, err := c.FormFile("image")
    if err == nil { // If a file is uploaded
        filename := filepath.Base(file.Filename)
        savePath := filepath.Join("static", "images", "categories", filename)

        // Save the uploaded file
        if saveErr := c.SaveUploadedFile(file, savePath); saveErr != nil {
            c.HTML(http.StatusInternalServerError, "category_edit.html", gin.H{
                "error": "Failed to upload image",
                "category": category,
            })
            return
        }

        // Update the image path in the database
        category.ImagePath = filepath.ToSlash(savePath) // Normalize path to use forward slashes
    }

    // Save the updated category
    if err := db.Save(&category).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "category_edit.html", gin.H{
            "error": "Failed to update category",
            "category": category,
        })
        return
    }

    // Redirect to the categories list
    c.Redirect(http.StatusFound, "/categories")
}


func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := db.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.Redirect(http.StatusFound, "/categories")
}
