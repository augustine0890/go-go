package main

import (

	// "html/template"
	"fmt"
	"log"
	"strconv"

	"snippetbox/pkg/models/mysql"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	repo mysql.Repository
}

func NewService(r mysql.Repository) service {
	return service{
		repo: r,
	}
}

func home() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("home.page", fiber.Map{})
	}
}

func showSnippet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 0 {
		return c.SendStatus(400)
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	return c.SendString(msg)
}

func createSnippet(s service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		title := "0 snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := "7"

		id, err := s.repo.Insert(title, content, expires)
		log.Println(err)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return ctx.Redirect(fmt.Sprintf("/snippet?id=%d", id), 301)
	}
}
