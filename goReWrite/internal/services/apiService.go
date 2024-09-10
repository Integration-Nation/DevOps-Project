package services

import (
	"DevOps-Project/internal/models"
)

func GetSearchResults(query string, language string) []models.Page {
 if query == "" {
        return []Page{} // Return empty if no query
    }
    
    return repository.SearchInDB(query, language)
}