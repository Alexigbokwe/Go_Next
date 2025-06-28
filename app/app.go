package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App
}

func NewApp() *App {
	return &App{
		App: fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			},
		}),
	}
}

func (a *App) Listen(addr string) {
	//
	a.App.Use(recover.New())

	// Default GET route
	a.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to GoNext framework")
	})

	if err := a.App.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Called when a module is initialized.
func (app *App) InitModules(modules []Module, container *Container) {
	for _, module := range modules {
		if hook, ok := module.(OnModuleInit); ok {
			if err := hook.OnModuleInit(); err != nil {
				log.Fatal(err)
			}
		}
		module.Register(container)
		module.MountRoutes(app)
	}
}

func (app *App) ShutdownModules(modules []Module) {
	for _, module := range modules {
		if hook, ok := module.(OnModuleDestroy); ok {
			if err := hook.OnModuleDestroy(); err != nil {
				log.Println("Error during shutdown:", err)
			}
		}
	}
}
