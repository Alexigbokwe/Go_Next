# GoNext Framework

A scalable, modular Go web application framework using [Fiber](https://gofiber.io/), with built-in dependency injection (DI) supporting Singleton, Scoped, and Transient service lifetimes. Designed for rapid development of maintainable, testable, and production-ready Go web services.

---

## Features

- **Fiber v2**: Fast, Express-inspired web framework for Go.
- **Modular Structure**: Organize code by feature/module for maintainability.
- **Dependency Injection**: Custom DI container with support for:
  - **Singleton**: One instance for the app lifetime (default).
  - **Scoped**: One instance per HTTP request.
  - **Transient**: New instance every time it's resolved.
- **Global Middleware**: Register middleware that applies to all routes in one place.
- **Guards**: Protect routes with custom logic (e.g., role checks).
- **DTOs & Validation**: Use Data Transfer Objects and struct validation for request data.
- **Module Lifecycle Hooks**: Modules can run code on init and destroy.
- **Global Error Handling**: Prevents server crashes and returns clean JSON errors.
- **Request-Scoped Services**: Easily access per-request data (e.g., User-Agent).
- **Ready for Testing**: Decoupled components for easy unit and integration testing.

---

## Project Structure

```
.
├── app/                # Core application, DI container, module registration, guards, middleware, validation
│   ├── app.go          # App struct, lifecycle, error handling
│   ├── container.go
│   ├── fiber.go
│   ├── guard.go        # Guard interface
│   ├── middleware.go   # Middleware interface
│   ├── module.go
│   ├── registerModuleComponents.go
│   └── validator.go    # Validation helper
├── global/             # Global middleware (e.g., logging)
│   └── globalMiddleware/
│       └── logging.go
├── internal/           # Application modules (feature-based)
│   └── user/           # Example: User management module
│       ├── controller/
│       │   └── userController.go
│       ├── dto/        # Data Transfer Objects for validation
│       │   └── createUserDto.go
│       ├── guard/      # Guards for route protection
│       │   └── adminGuard.go
│       ├── repository/
│       │   └── userRepository.go
│       ├── route/
│       │   └── userRoute.go
│       ├── service/
│       │   └── userService.go
│       └── module.go
├── main.go             # Application entry point
├── go.mod
└── go.sum
```

---

## Getting Started

### Prerequisites

- Go 1.18+
- [Fiber v2](https://gofiber.io/)

### Installation

Use gonext CLI to start a new project. First install gonext CLI

```bash
go install github.com/Alexigbokwe/gonext@latest
```

Once gonext CLI is installed, you can now start a new project with the below command.

```bash
gonext new <project_name>
```

---

## Usage

### 1. **Creating a Module**

Each feature (e.g., users, products) is a module under `internal/`.  
A module typically contains:

- `controller/` — HTTP handlers
- `service/` — Business logic
- `repository/` — Data access
- `route/` — Route registration
- `dto/` — Data Transfer Objects for request validation
- `guard/` — Guards for route protection
- `module.go` — Module registration

### 2. **App Initialization & Global Middleware**

Global middleware is registered in `main.go`:

```go
func registerGlobalMiddleware(app *goNext.App) {
    app.Use(globalMiddleware.LoggingMiddleware{}.Use())
}

func main() {
    app := goNext.NewApp()
    container := goNext.NewContainer()
    registerGlobalMiddleware(app)
    modules := registerModules()
    app.InitModules(modules, container)
    // ...
}
```

### 3. **Module Lifecycle Hooks**

Modules can implement `OnModuleInit` and `OnModuleDestroy` for setup/teardown:

```go
type UserModule struct {}

func (m *UserModule) OnModuleInit() error {
    fmt.Println("UserModule initialized!")
    return nil
}

func (m *UserModule) OnModuleDestroy() error {
    fmt.Println("UserModule destroyed!")
    return nil
}
```

### 4. **Guards: Route Protection**

Guards allow you to protect routes with custom logic:

```go
// app/guard.go
// interface Guard { CanActivate(ctx *fiber.Ctx) bool }

// internal/user/guard/adminGuard.go
type AdminGuard struct{}
func (g AdminGuard) CanActivate(ctx *fiber.Ctx) bool {
    return ctx.Get("X-Admin") == "true"
}

// Guard middleware helper
func GuardMiddleware(g app.Guard) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if !g.CanActivate(c) {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden: Guard rejected request"})
        }
        return c.Next()
    }
}

// Usage in route
group.Post("/users", guard.GuardMiddleware(guard.AdminGuard{}), ctrl.CreateUser)
```

### 5. **DTOs & Validation**

Use DTOs for request validation:

```go
// internal/user/dto/createUserDto.go
type CreateUserDTO struct {
    Name  string `json:"name" validate:"required,min=2"`
    Email string `json:"email" validate:"required,email"`
}

// In controller
var dto dto.CreateUserDTO
if err := c.BodyParser(&dto); err != nil {
    return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
}
if err := app.ValidateStruct(dto); err != nil {
    return c.Status(422).JSON(fiber.Map{"validation": err.Error()})
}
```

### 6. **Module Registration Example**

```go
import userModule "goNext/internal/user"

modules := []app.Module{
    userModule.NewUserModule(),
}
```

---

## Global Error Handling

Fiber is configured with a global error handler and panic recovery:

```go
app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    },
})
app.Use(recover.New())
```

---

## Adding a New Module

1. Create a new folder under `internal/` (e.g., `productMs`).
2. Add `controller/`, `service/`, `repository/`, `route/`, `dto/`, `guard/`, and `module.go`.
3. Register your module in `main.go` as shown above.

---

## Testing

- Write unit tests for services and repositories.
- Use dependency injection to mock dependencies.

---

## Extending the Template

- Add authentication/authorization modules using guards.
- Integrate with databases (Postgres, MySQL, etc.).
- Add middleware (logging, CORS, etc.).
- Add request validation and response formatting.

---

## Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes
4. Push to the branch (`git push origin feature/fooBar`)
5. Open a pull request

---

## License

MIT

---

**Happy coding!**  
This framework is designed to help you build robust, maintainable Go web applications with best practices from day one.
