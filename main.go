package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() //create a Gin router with default middleware(logging and recovery)
	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok!!",
		})
	})
	server.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin")
	})

	//dynamic route
	server.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id") //extract id from url
		c.String(200, "User ID: %v", id)
	})

	server.Run(":8080") //listening on port 8080
}
