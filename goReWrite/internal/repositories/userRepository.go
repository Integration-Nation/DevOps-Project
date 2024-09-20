package repositories

import (
	"DevOps-Project/internal/models"

	"gorm.io/gorm"
)

type UserRepositoryI interface {
	FindByUsername(username string) (*models.User, error)
	GetAllUsers() *[]models.User
	CreateUser(username, email, password string) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
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
        return nil, err
    }

    // The user.ID is now populated after Create()
    return &user, nil
}