package repositories

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/utilities"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepositoryI interface {
	FindByUsername(username string) (*models.User, error)
	GetAllUsers() *[]models.User
	CreateUser(username, email, password string) (*models.User, error)
	DeleteUser(username string) error
}

type UserRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: utilities.NewLogger(),
	}

}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		ur.logger.Error("Error while fetching user by username", zap.Error(result.Error))
		return nil, result.Error
	}
	return &user, nil

}

func (ur *UserRepository) GetAllUsers() *[]models.User {
	var users []models.User

	ur.db.Find(&users)

	return &users

}

func (ur *UserRepository) CreateUser(username, email, password string) (*models.User, error) {
	user := models.User{
		Username: username,
		Email:    email,
		Password: password, // Storing password as plain text for now (not recommended in production)
	}

	// Create the user in the database
	err := ur.db.Create(&user).Error
	if err != nil {
		ur.logger.Error("Error while creating user", zap.Error(err))
		return nil, err
	}

	// The user.ID is now populated after Create()
	ur.logger.Info("User created successfully", zap.String("username", user.Username))
	return &user, nil
}

// delete user by username
func (ur *UserRepository) DeleteUser(username string) error {
	var user models.User

	result := ur.db.Where("username = ?", username).First(&user).Delete(&user)

	if result.Error != nil {
		ur.logger.Error("Error while deleting user by username", zap.Error(result.Error))
		return result.Error
	}
	return nil
}
