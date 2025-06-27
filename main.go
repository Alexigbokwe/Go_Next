package main

import (
	goNext "goNext/app"
	userModule "goNext/internal/user"
)

func main() {
	app := goNext.NewApp()
	container := goNext.NewContainer()

	// plug in modules here
	modules := []goNext.Module{
		userModule.NewUserModule(),
	}

	// Register and wire them
	for _, module := range modules {
		module.Register(container)
		module.MountRoutes(app)
	}

	app.Listen(":5050")
}
