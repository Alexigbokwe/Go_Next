package user

import (
	"fmt"
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

// Called when a module is initialized.
func (m *UserModule) OnModuleInit() error {
	fmt.Println("UserModule initialized!")
	return nil
}

// Called when a module is destroyed.
func (m *UserModule) OnModuleDestroy() error {
	fmt.Println("UserModule destroyed!")
	return nil
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
