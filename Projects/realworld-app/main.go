package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Start the default gin server
	r := gin.Default()

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please create one in the root directory")
	}

	r.LoadHTMLGlob("./public/html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"realworldVersion": "v1.01",
			"goVersion":        runtime.Version(),
		})
	})

	v1 := r.Group("/api/v1")
	{
		user := new()
	}
	port := os.Getenv("PORT")

	r.Run(":" + port)
}
