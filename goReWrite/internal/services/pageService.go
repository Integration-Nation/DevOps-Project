package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"

	"go.uber.org/zap"
)

type PageServiceI interface {
	GetSearchResults(q string, language string) ([]models.Page, error)
}

type PageService struct {
	repo   repositories.PageRepositoryI
	logger *zap.Logger
}

func NewPageService(repo repositories.PageRepositoryI, logger *zap.Logger) *PageService {
	return &PageService{
		repo:   repo,
		logger: logger,
	}
}

func (ps *PageService) GetSearchResults(q string, language string) ([]models.Page, error) {
	return ps.repo.GetSearchResults(q, language)
}
