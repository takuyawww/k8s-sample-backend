package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "app-1",
			"version": "0.0.2",
		})
	})
	r.Run(":8081")
}
