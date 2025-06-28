package guard

import (
	"goNext/app"

	"github.com/gofiber/fiber/v2"
)

type AdminGuard struct{}

func (g AdminGuard) CanActivate(ctx *fiber.Ctx) bool {
	// Dummy check: In real scenario, verify JWT or user roles
	return ctx.Get("X-Admin") == "true"
}

// Helper to wrap guards
func GuardMiddleware(g app.Guard) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !g.CanActivate(c) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Guard rejected request",
			})
		}
		return c.Next()
	}
}
