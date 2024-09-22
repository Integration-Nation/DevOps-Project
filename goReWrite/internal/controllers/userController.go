package controllers

import (
	"DevOps-Project/internal/models"
	"DevOps-Project/internal/services"
	"DevOps-Project/internal/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserControllerI interface {
	Login(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type UserController struct {
	service  services.UserServiceI
	validate *validator.Validate
	logger   *zap.Logger
}

func NewUserController(service services.UserServiceI, validate *validator.Validate) *UserController {
	return &UserController{
		service:  service,
		validate: validate,
		logger:   utilities.NewLogger(),
	}
}

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

	// Verify the user credentials
	user, err := uc.service.VerifyLogin(req.Username, req.Password)
	if err != nil {
		// Return an unauthorized status with an error message in JSON
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// On success, return a JSON response with the user info or token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Login successful",
		"user_id":  user.ID,
		"username": user.Username,
	})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {

	// Get all users from the database
	users := uc.service.GetAllUsers()

	// On success, return a JSON response with the users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}

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

func (uc *UserController) Logout(c *fiber.Ctx) error {
	return c.SendString("Logged Out")
}
