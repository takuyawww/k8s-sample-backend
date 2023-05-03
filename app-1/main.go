package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "app-1",
			"version": "1.0.0",
		})
	})
	r.Run(":8081")
}
