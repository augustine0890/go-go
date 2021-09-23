package main

import (
	"errors"
	"fmt"
	"snippetbox/pkg/models"
	"snippetbox/pkg/models/mysql"
	"strconv"

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

func home(s service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s, err := s.repo.Latest()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return ctx.Status(fiber.StatusOK).JSON(s)
		// return ctx.Render("home.page", fiber.Map{})
	}
}

func showSnippet(s service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil || id < 0 {
			return ctx.SendStatus(400)
		}
		s, err := s.repo.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				return ctx.Status(fiber.StatusNotFound).SendString("Record not found")
			} else {
				return err
			}
		}

		return ctx.Render("show.page", &fiber.Map{
			"snippet": s,
		})
		// return ctx.Status(fiber.StatusOK).JSON(s)
	}
}

func createSnippet(s service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		title := "0 snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := "7"

		id, err := s.repo.Insert(title, content, expires)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return ctx.Redirect(fmt.Sprintf("/snippet?id=%d", id), 301)
	}
}
