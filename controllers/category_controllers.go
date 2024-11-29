package controllers
import (
	"go-final-project/models"
	"go-final-project/config"
	"github.com/gin-gonic/gin"
	"net/http"
	
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


