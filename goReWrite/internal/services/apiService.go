package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
)

type PageService interface {
	SearchPages(q string, language string) ([]models.Page, error)
}

type pageService struct {
	repo repositories.PageRepository
}

func NewPageService(repo repository.PageRepository) PageService {
	return &pageService{repo}
}

func GetSearchResults(q string, language string) ([]models.Page, error) {
	return repositories.GetSearchResults(q, language)
}
