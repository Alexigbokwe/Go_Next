package globalMiddleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type LoggingMiddleware struct{}

func (m LoggingMiddleware) Use() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("Incoming request: %s %s", c.Method(), c.Path())
		return c.Next()
	}
}
