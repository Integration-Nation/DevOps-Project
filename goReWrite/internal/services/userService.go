package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/monitoring"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/security"
	"errors"
	"time"

	"go.uber.org/zap"
)

type UserServiceI interface {
	VerifyLogin(username, password string) (string, string, error)
	GetAllUsers() *[]models.User
	RegisterUser(username, email, password string) (*models.User, error)
	DeleteUser(username string) (string, error)
	collectUserMetrics()
}

type UserService struct {
	repo   repositories.UserRepositoryI
	logger *zap.Logger
}

func NewUserService(repo repositories.UserRepositoryI, logger *zap.Logger) *UserService {
	service := &UserService{
		repo:   repo,
		logger: logger,
	}

	go service.collectUserMetrics()
	return service
}

func (us *UserService) VerifyLogin(username, password string) (string, string, error) {
	user, err := us.repo.FindByUsername(username)
	if err != nil {
		us.logger.Error("Failed to find user by username", zap.Error(err))
		return "", "", errors.New("internal server error")
	}

	if err := security.ComparePasswords(user.Password, password); err != nil {
		us.logger.Error("Failed to compare passwords", zap.Error(err))
		return "", "", errors.New("internal server error")
	}

	token, err := security.GenerateJWT(user.ID, user.Username)
	if err != nil {
		us.logger.Error("Failed to generate JWT token", zap.Error(err))
		return "", "", errors.New("could not generate token")
	}

	return token, user.Username, nil

}

func (us *UserService) GetAllUsers() *[]models.User {
	return us.repo.GetAllUsers()
}

func (us *UserService) RegisterUser(username, email, password string) (*models.User, error) {

	// Check if username already exists
	existingUser, _ := us.repo.FindByUsername(username)

	if existingUser != nil {
		us.logger.Info("Username already taken", zap.String("username", username))
		return nil, errors.New("the username is already taken")
	}

	hashedPassword, err := security.HashPassword(password)
	if err != nil {
		us.logger.Error("Failed to hash password", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	// Create the user
	return us.repo.CreateUser(username, email, hashedPassword)
}

func (us *UserService) DeleteUser(username string) (string, error) {
	err := us.repo.DeleteUser(username)
	if err != nil {
		us.logger.Error("Failed to delete user", zap.Error(err))
		return "", errors.New("internal server error")
	}

	return username + " has been deleted", nil
}

func (us *UserService) collectUserMetrics() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		// Update total users
		totalUsers, err := us.repo.CountTotal()
		if err == nil {
			monitoring.UpdateTotalUsers(totalUsers)
		}

		// Update active users for different periods
		now := time.Now()

		// Daily active users (last 24 hours)
		dailyActive, err := us.repo.CountActiveUsers(now.Add(-24 * time.Hour))
		if err == nil {
			monitoring.UpdateActiveUsers("daily", dailyActive)
		}

		// Weekly active users (last 7 days)
		weeklyActive, err := us.repo.CountActiveUsers(now.Add(-7 * 24 * time.Hour))
		if err == nil {
			monitoring.UpdateActiveUsers("weekly", weeklyActive)
		}

		// Monthly active users (last 30 days)
		monthlyActive, err := us.repo.CountActiveUsers(now.Add(-30 * 24 * time.Hour))
		if err == nil {
			monitoring.UpdateActiveUsers("monthly", monthlyActive)
		}
	}
}
