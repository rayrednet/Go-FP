package main

import (
	"github.com/gin-gonic/gin"
	"go-final-project/config" 
)

func main() {
	
	db := config.GetDB() 
	r := gin.Default()

	
	r.LoadHTMLGlob("templates/*")

	
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	r.Run(":8080")
}
