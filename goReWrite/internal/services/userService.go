package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/utilities"
	"errors"

	"go.uber.org/zap"
)

type UserServiceI interface {
	VerifyLogin(username, password string) (*models.User, error)
	GetAllUsers() *[]models.User
	RegisterUser(username, email, password string) (*models.User, error)
}

type UserService struct {
	repo   repositories.UserRepositoryI
	logger *zap.Logger
}

func NewUserService(repo repositories.UserRepositoryI) *UserService {
	return &UserService{
		repo:   repo,
		logger: utilities.NewLogger(),
	}
}

func (us *UserService) VerifyLogin(username, password string) (*models.User, error) {
	return us.repo.FindByUsername(username)

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

	// Create the user
	return us.repo.CreateUser(username, email, password)
}
