package repositories

import "DevOps-Project/internal/models"

func GetSearchResults(query string, language string) []models.Page {
	 var pages []models.Page

    DB.Where("language = ? AND content LIKE ?", language, "%"+query+"%").Find(&pages)

    return pages
}