package main

import "github.com/gin-gonic/gin"

fun main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
					"message": "Hello, Gin!",
			})
	})

	r.Run(":8080")

}