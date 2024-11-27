package controllers

import (
	"DevOps-Project/internal/services"

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

func NewPageController(service services.PageServiceI, validate *validator.Validate, logger *zap.Logger) *PageController {
	return &PageController{
		service:  service,
		validate: validate,
		logger:   logger,
	}
}

// GetSearchResults godoc
// @Summary Get search results
// @Description Get search results from the service based on a query and language parameter
// @Tags search
// @Produce json
// @Param q query string true "Search query"
// @Param language query string false "Language of the results, default is 'en'" default(en)
// @Success 200 {object} map[string]interface{} "List of search results"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /search [get]
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
