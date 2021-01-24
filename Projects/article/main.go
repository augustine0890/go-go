package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Load the templates
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start serving the application
	router.Run()
}
