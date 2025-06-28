/**
 * Main Application Entry Point
 *
 * This file serves as the entry point for the Go application. It initializes
 * the application, registers global middleware, and wires up modules to ensure
 * proper functionality and separation of concerns.
 *
 * Key Functions:
 * - registerGlobalMiddleware: Registers middleware that applies globally to all routes.
 * - registerModules: Returns a slice of modules to be registered in the application.
 * - main: Initializes the application, sets up the container, registers middleware and modules,
 *         and starts the server.
 *
 * Usage:
 * - Define global middleware in the `registerGlobalMiddleware` function.
 * - Add new modules to the `registerModules` function to extend application functionality.
 * - Ensure each module implements the `goNext.Module` interface for proper integration.
 *
 * Port:
 * - The application listens on port `5050` by default.
 *
 * Dependencies:
 * - goNext/app: Core application framework.
 * - goNext/global/globalMiddleware: Middleware for global application behavior.
 */
package main

import (
	goNext "goNext/app"
	globalMiddleware "goNext/global/globalMiddleware"
	userModule "goNext/internal/user"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func registerGlobalMiddleware(app *goNext.App) {
	app.Use(globalMiddleware.LoggingMiddleware{}.Use())
}

func registerModules() []goNext.Module {
	return []goNext.Module{
		userModule.NewUserModule(),
	}
}

func main() {
	app := goNext.NewApp()
	container := goNext.NewContainer()
	registerGlobalMiddleware(app)
	modules := registerModules()

	app.InitModules(modules, container)

	go func() {
		app.Listen(":5050")
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Call shutdown hooks
	app.ShutdownModules(modules)
	log.Println("Server gracefully stopped")
}
