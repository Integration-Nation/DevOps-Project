package controllers

import (
	"DevOps-Project/internal/services"
	"DevOps-Project/internal/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type PageControllerI interface {
	GetSearchResults(c *fiber.Ctx) error
}

type PageController struct {
	service  services.PageServiceI
	validate *validator.Validate
	logger   *zap.Logger
}

func NewPageController(service services.PageServiceI, validate *validator.Validate) *PageController {
	return &PageController{
		service:  service,
		validate: validate,
		logger:   utilities.NewLogger(),
	}
}

// Ctx er res og req
func (pc *PageController) GetSearchResults(c *fiber.Ctx) error {
	q := c.Query("q")
	language := c.Query("language", "en")

	if q == "" {
		return c.JSON(fiber.Map{"search_results": []string{}})
	}

	pages, err := pc.service.GetSearchResults(q, language)
	if err != nil {
		pc.logger.Error("Error getting search results", zap.String("error", err.Error()), zap.String("query", q))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"search_results": pages})
}
