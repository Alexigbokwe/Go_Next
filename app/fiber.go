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
	a.App.Use(recover.New())
	if err := a.App.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
