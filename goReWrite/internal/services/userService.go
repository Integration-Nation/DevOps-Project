package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
	"DevOps-Project/internal/security"
	"DevOps-Project/internal/utilities"
	"errors"

	"go.uber.org/zap"
)

type UserServiceI interface {
	VerifyLogin(username, password string) (string, string, error)
	GetAllUsers() *[]models.User
	RegisterUser(username, email, password string) (*models.User, error)
	DeleteUser(username string) (string, error)
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

	token, err := security.GenerateJWT(int(user.ID), user.Username)
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
