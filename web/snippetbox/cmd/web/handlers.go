package main

import (

	// "html/template"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func home(c *fiber.Ctx) error {
	return c.Render("home.page", fiber.Map{})
}

func showSnippet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 0 {
		return c.SendStatus(400)
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	return c.SendString(msg)
}

func createSnippet(c *fiber.Ctx) error {
	return c.SendString("Create a new snippet...")
}
