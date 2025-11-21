package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
			"expires_in": 3600,
		})
	})

	r.GET("/verify", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"valid": true})
	})

	r.Run(":8081")
}
