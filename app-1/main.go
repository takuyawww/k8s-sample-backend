package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/version", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "app-1",
			"version": "0.0.3",
		})
	})
	r.Run(":8081")
}
