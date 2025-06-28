package route

import (
	"goNext/internal/user/controller"
	"goNext/internal/user/guard"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(route fiber.Router, ctrl *controller.UserController) {
	route.Get("/", ctrl.GetUsers)

	route.Post("/users", guard.GuardMiddleware(guard.AdminGuard{}), ctrl.CreateUser)
}
