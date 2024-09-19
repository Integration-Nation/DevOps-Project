package repositories

import (
	"DevOps-Project/internal/initializers"
	"DevOps-Project/internal/models"

	"gorm.io/gorm"
)

type PageRepositoryI interface {
	GetSearchResults(q string, language string) ([]models.Page, error)
}

type PageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{db: db}
}

func (pr *PageRepository) GetSearchResults(q string, language string) ([]models.Page, error) {
	var pages []models.Page
	query := "%" + q + "%"
	err := initializers.DB.Where("language = ? AND content LIKE ?", language, query).Find(&pages).Error
	return pages, err
}
