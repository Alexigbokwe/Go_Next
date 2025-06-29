package controller

import (
	"goNext/app"
	"goNext/internal/user/dto"
	"goNext/internal/user/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *service.UserService `inject:"type"`
}

// GetUsers handles the request to get users
func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	result := uc.UserService.GetUsers()
	if !result.Status {
		return c.Status(result.Code).JSON(result)
	}
	return c.Status(result.Code).JSON(result)
}

// CreateUser handles the request to create a user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var dto dto.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := app.ValidateStruct(dto); err != nil {
		return c.Status(422).JSON(fiber.Map{"validation": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User created"})
}

// UpdateUser handles the request to update a user
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	// Implementation for updating a user
	return nil
}

// DeleteUser handles the request to delete a user
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	// Implementation for deleting a user
	return nil
}
