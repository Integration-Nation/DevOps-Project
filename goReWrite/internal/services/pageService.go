package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/utilities"

	"go.uber.org/zap"
)

type PageServiceI interface {
	GetSearchResults(q string, language string) ([]models.Page, error)
}

type PageService struct {
	repo   repositories.PageRepositoryI
	logger *zap.Logger
}

func NewPageService(repo repositories.PageRepositoryI) *PageService {
	return &PageService{
		repo:   repo,
		logger: utilities.NewLogger(),
	}
}

func (ps *PageService) GetSearchResults(q string, language string) ([]models.Page, error) {
	return ps.repo.GetSearchResults(q, language)
}
