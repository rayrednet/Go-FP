package controllers

import "github.com/gin-gonic/gin"

func HomeHandler(c *gin.Context) {
	c.HTML(200, "layout.html", gin.H{
		"title": "Home - Scarlett's Café",
	})
}
