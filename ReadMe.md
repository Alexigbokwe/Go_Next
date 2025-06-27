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
- **Global Error Handling**: Prevents server crashes and returns clean JSON errors.
- **Request-Scoped Services**: Easily access per-request data (e.g., User-Agent).
- **Ready for Testing**: Decoupled components for easy unit and integration testing.

---

## Project Structure

```
.
├── app/                # Core application, DI container, module registration
│   ├── container.go
│   ├── fiber.go
│   ├── module.go
│   └── registerModuleComponents.go
├── internal/           # Application modules (feature-based)
│   └── user/         # Example: User management module
│       ├── controller/
│       │   └── userController.go
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
go install github.com/Alexigbokwe/goNext_CLI@latest
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
- `module.go` — Module registration

### Start the Project

Start your GoNext project:

```sh
gonext start
```

Start your project in watch mode (hot reload):

```sh
gonext start --watch
```

> **Note:** Watch mode requires [`air`](https://github.com/cosmtrek/air). Install it with:
>
> ```sh
> go install github.com/cosmtrek/air@latest
> ```

### Generate Modules and Components

- Generate a new module:

  ```sh
  gonext generate module <name>
  # or
  gonext g module <name>
  ```

- Generate a controller, service, or repository in a module (creates the module if it doesn't exist):
  ```sh
  gonext generate controller <name> <in_module>
  gonext generate service <name> <in_module>
  gonext generate repository <name> <in_module>
  # or use the 'g' alias
  gonext g controller <name> <in_module>
  gonext g service <name> <in_module>
  gonext g repository <name> <in_module>
  ```

### 2. **Dependency Injection & Service Scopes**

#### **Singleton (default)**

- One instance for the entire app.
- Use for stateless services, repositories, controllers.

#### **Scoped**

- One instance per HTTP request.
- Use for request-specific data (e.g., current user, request context).

#### **Transient**

- New instance every time it's resolved.
- Use for short-lived, stateless objects.

**Example: Registering Services with Scopes**

```go
// Singleton (default)
container.Register(&MySingletonService{})

// Scoped
container.RegisterScoped(func(ctx *fiber.Ctx) *MyScopedService {
    return NewMyScopedService(ctx)
})

// Transient
container.RegisterTransient(func() *MyTransientService {
    return NewMyTransientService()
})
```

### 3. **Accessing Request Data with Scoped Services**

**Example: Get browser name from User-Agent header per request**

```go
// service/browserInfoService.go
type BrowserInfoService struct {
    BrowserName string
}
func NewBrowserInfoService(c *fiber.Ctx) *BrowserInfoService {
    return &BrowserInfoService{BrowserName: c.Get("User-Agent")}
}

// Register as scoped
container.RegisterScoped(func(c *fiber.Ctx) *BrowserInfoService {
    return NewBrowserInfoService(c)
})

// Inject into controller
type UserController struct {
    BrowserInfoService *service.BrowserInfoService `inject:"type"`
}
```

---

## Global Error Handling

Fiber is configured with a global error handler and panic recovery:

```go
import "github.com/gofiber/fiber/v2/middleware/recover"

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
2. Add `controller/`, `service/`, `repository/`, `route/`, and `module.go`.
3. Register your module in `main.go`:

```go
import productModule "goNext/internal/productMs"

modules := []app.Module{
    userModule.NewUserModule(),
    productModule.NewProductModule(),
}
```

---

## Testing

- Write unit tests for services and repositories.
- Use dependency injection to mock dependencies.

---

## Extending the Template

- Add authentication/authorization modules.
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
