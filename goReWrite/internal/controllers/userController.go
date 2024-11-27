package controllers

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserControllerI interface {
	Login(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserController struct {
	service  services.UserServiceI
	validate *validator.Validate
	logger   *zap.Logger
}

func NewUserController(service services.UserServiceI, validate *validator.Validate, logger *zap.Logger) *UserController {
	return &UserController{
		service:  service,
		validate: validate,
		logger:   logger,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user with username and password, and return a token
// @Tags users
// @Accept json
// @Produce json
// @Param loginRequest body models.LoginRequest true "Login request"
// @Success 200 {object} map[string]interface{} "Token and username"
// @Failure 400 {object} map[string]interface{} "Invalid request format or validation errors"
// @Failure 401 {object} map[string]interface{} "Unauthorized, wrong credentials"
// @Router /users/login [post]
func (uc *UserController) Login(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		uc.logger.Error("Error parsing login request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := uc.validate.Struct(req); err != nil {
		uc.logger.Error("Error validating login request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, username, err := uc.service.VerifyLogin(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":    token,
		"username": username,
	})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all registered users
// @Tags users
// @Produce json
// @Success 200 {object} map[string]interface{} "List of users"
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users := uc.service.GetAllUsers()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param registerRequest body models.RegisterRequest true "Register request"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or validation errors"
// @Router /users/register [post]
func (uc *UserController) Register(c *fiber.Ctx) error {
	var req models.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		uc.logger.Error("Error parsing register request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := uc.validate.Struct(req); err != nil {
		uc.logger.Error("Error validating register request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := uc.service.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		uc.logger.Error("Error registering user", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully!",
	})
}

// Logout godoc
// @Summary Logout user
// @Description Logout the authenticated user
// @Tags users
// @Success 200 {string} string "Logged Out"
// @Router /users/logout [post]
func (uc *UserController) Logout(c *fiber.Ctx) error {
	return c.SendString("Logged Out")
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by username
// @Tags users
// @Param username query string true "Username of the user to delete"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid username"
// @Router /users [delete]
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	username := c.Query("username")
	response, err := uc.service.DeleteUser(username)
	if err != nil {
		return err
	}
	return c.SendString(response)
}
