package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []string     `json:"steps"`
	Picture     string       `json:"imageURL"`
}

type Ingredient struct {
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

var recipes []Recipe

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"recipes": recipes,
	})
}

func RecipeHandler(c *gin.Context) {
	for _, recipe := range recipes {
		if recipe.ID == c.Param("id") {
			c.HTML(http.StatusOK, "recipe.html", gin.H{
				"recipe": recipe,
			})
			return
		}
	}
	c.File("404.html")
}

// func init() {
// recipes = make([]Recipe, 0)
// data, _ := f.ReadFile("recipes.json")
// json.Unmarshal(data, &recipes)
// }

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {

	router := gin.Default()
	router.Static("/assets", "./assets")
	// router.StaticFS("/assets", http.FS(fsys))
	router.LoadHTMLGlob("templates/*.html")
	// router.SetHTMLTemplate(templ)

	router.GET("/", IndexHandler)
	router.GET("/recipes/:id", RecipeHandler)
	router.Run()
	fmt.Println("recipes", recipes)

}
