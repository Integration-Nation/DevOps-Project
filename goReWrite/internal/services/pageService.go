package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
)

type PageServiceI interface {
	GetSearchResults(q string, language string) ([]models.Page, error)
}

type PageService struct {
	repo repositories.PageRepositoryI
}

func NewPageService(repo repositories.PageRepositoryI) *PageService {
	return &PageService{repo: repo}
}

func (ps *PageService) GetSearchResults(q string, language string) ([]models.Page, error) {
	return ps.repo.GetSearchResults(q, language)
}
