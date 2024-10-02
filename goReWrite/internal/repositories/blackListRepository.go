package repositories

import (
	"DevOps-Project/internal/models"
	"time"

	"gorm.io/gorm"
)

type TokenBlacklistRepositoryI interface {
	AddToken(token string, duration time.Duration) error
	IsBlacklisted(token string) (bool, error)
	CleanupExpiredTokens() error
}

type TokenBlacklistRepository struct {
	db *gorm.DB
}

func NewTokenBlacklistRepository(db *gorm.DB) *TokenBlacklistRepository {
	return &TokenBlacklistRepository{db: db}
}

func (br *TokenBlacklistRepository) AddToken(token string, duration time.Duration) error {
	expiresAt := time.Now().Add(duration)
	blacklistEntry := models.TokenBlacklist{Token: token, ExpiresAt: expiresAt}
	return br.db.Create(&blacklistEntry).Error
}

func (br *TokenBlacklistRepository) IsBlacklisted(token string) (bool, error) {
	var count int64
	err := br.db.Model(&models.TokenBlacklist{}).Where("token = ? AND expires_at > ?", token, time.Now()).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (br *TokenBlacklistRepository) CleanupExpiredTokens() error {
	return br.db.Where("expires_at < ?", time.Now()).Delete(&models.TokenBlacklist{}).Error
}
