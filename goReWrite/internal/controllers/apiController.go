package controllers

import (
	"DevOps-Project/internal/services"

	"github.com/gofiber/fiber/v2"
)

type PageController struct {
	service services.PageService
}

func NewPageController(service services.PageService) *PageController {
	return &PageController{service}
}

// Ctx er res og req
func GetSearchResults(c *fiber.Ctx) error {
	q := c.Query("q")
	language := c.Query("language", "en")

	if q == "" {
		return c.JSON(fiber.Map{"search_results": []string{}})
	}

	pages, err := services.GetSearchResults(q, language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"search_results": pages})
}
