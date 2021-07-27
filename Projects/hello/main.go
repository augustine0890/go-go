package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Person struct {
	XMLName xml.Name `xml:"person"`
	FirstName string `xml:"firstName,attr"`
	LastName string `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(http.StatusOK, Person{
		FirstName: "Augustine",
		LastName: "Nguyen",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":8080")
}
