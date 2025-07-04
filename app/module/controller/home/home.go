package home

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Davion Asset Management Service using Go is running normally.",
	})
}
