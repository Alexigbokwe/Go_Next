package app

import "github.com/gofiber/fiber/v2"

type Middleware interface {
	Use() fiber.Handler
}
