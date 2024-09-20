package services

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/repositories"
)

type UserServiceI interface {
	VerifyLogin(username, password string) (*models.User, error)
	GetAllUsers() *[]models.User
}

type UserService struct {
	repo repositories.UserRepositoryI
}

func NewUserService(repo repositories.UserRepositoryI) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) VerifyLogin(username, password string) (*models.User, error) {
	return us.repo.FindByUsername(username)

}

func (us *UserService) GetAllUsers() *[]models.User {
	return us.repo.GetAllUsers()
}
