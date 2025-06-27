package user

import (
	"goNext/app"
	"goNext/internal/user/controller"
	"goNext/internal/user/repository"
	"goNext/internal/user/route"
	"goNext/internal/user/service"

	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	UserController *controller.UserController
}

func NewUserModule() *UserModule {
	return &UserModule{}
}

func (m *UserModule) Register(container *app.Container) {
	userRepo := &repository.UserRepository{}
	userService := &service.UserService{}
	userController := &controller.UserController{}

	app.RegisterModuleComponents(container, userRepo, userService, userController)

	m.UserController = userController // Use the same instance that was autowired
}

func (m *UserModule) MountRoutes(router fiber.Router) {
	// Mount user routes
	group := router.Group("/users")
	route.RegisterUserRoutes(group, m.UserController)
}
