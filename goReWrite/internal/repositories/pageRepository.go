package repositories

import (
	"DevOps-Project/internal/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PageRepositoryI interface {
	GetSearchResults(q string, language string) ([]models.Page, error)
}

type PageRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPageRepository(db *gorm.DB, logger *zap.Logger) *PageRepository {

	return &PageRepository{
		db:     db,
		logger: logger,
	}
}

func (pr *PageRepository) GetSearchResults(q string, language string) ([]models.Page, error) {
	var pages []models.Page
	query := "%" + q + "%"
	err := pr.db.Where("language = ? AND content LIKE ?", language, query).Find(&pages).Error
	return pages, err
}
