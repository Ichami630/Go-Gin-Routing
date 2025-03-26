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

	server.Run(":8080") //listening on port 8080
}
