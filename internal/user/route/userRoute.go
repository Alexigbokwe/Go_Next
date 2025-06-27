package route

import (
	"goNext/internal/user/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(route fiber.Router, ctrl *controller.UserController) {
	route.Get("/", ctrl.GetUsers)
}
