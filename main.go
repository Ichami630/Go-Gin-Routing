package main

import (
	"github.com/gin-gonic/gin"
)

// middleware
func AuthMiddleware(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if apiKey != "secret123" {
		c.JSON(401, gin.H{"error": "unathorised"})
		c.Abort() //stop further execution
		return
	}

	c.Next()
}

func main() {
	server := gin.Default() //create a Gin router with default middleware(logging and recovery)
	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok!!",
		})
	})
	// server.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Welcome to Gin")
	// })

	//dynamic route
	server.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id") //extract id from url
		c.String(200, "User ID: %v", id)
	})

	//routing with query params
	server.GET("/search", func(c *gin.Context) {
		query := c.Query("q") //extract the query param
		c.String(200, "Search result for %v", query)

	})

	//post request handling
	server.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		c.JSON(200, gin.H{
			"name":  name,
			"email": email,
		})
	})

	//grouping routes
	api := server.Group("/api")
	{
		api.GET("/users", func(c *gin.Context) {
			c.String(200, "list of users")
		})
		api.GET("/products", func(c *gin.Context) {
			c.String(200, "list of products")
		})
	} //endpoint localhost:8080/api/users

	//catch all route (404 not found)
	server.NoRoute(func(c *gin.Context) {
		c.String(404, "404 this route can not be resolved on the server")
	})

	//middlewares
	//acts as an intermediary btn 2 applications,services,systems, facilitating their communication/interaction

	//it is useful in logging,authorisation,requestvalidation,rate limiting

	//apply authentication middleware to protected route
	protected := server.Group("/admin")
	protected.Use(AuthMiddleware)
	protected.GET("/dashboard", func(c *gin.Context) {
		c.String(200, "welcome to admin dashboard")
	})

	//serving html file
	server.Static("/assets", "./static/assets") //serve static files like css,js,images

	server.LoadHTMLGlob("static/*.html") //load html files from the static folder

	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	server.Run(":8080") //listening on port 8080
}
