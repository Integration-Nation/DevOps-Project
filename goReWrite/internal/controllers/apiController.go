package controllers

import (
	"DevOps-Project/internal/services"

	"github.com/gofiber/fiber/v2"
)
type ApiController struct{
	service 
}

//ctx er req og res
func GetSearchResults(c *fiber.Ctx)  error{
    // Get query params
    query := c.Query("q", "")
    language := c.Query("language", "en")

    // Call service layer
    searchResults := services.GetSearchResults(query, language)

    // Return the response
    return c.JSON(fiber.Map{
        "search_results": searchResults,
    })
}