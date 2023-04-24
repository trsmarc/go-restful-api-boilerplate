package router

import (
	"github.com/marktrs/simple-todo/handler"
	"github.com/marktrs/simple-todo/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Simple-TODO API Metrics"}))

	api := app.Group("/api", logger.New())
	api.Get("/health", handler.HealthCheck)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/users")
	user.Post("/", handler.CreateUser)

	// Task
	task := api.Group("/tasks")
	task.Use(middleware.Protected())
	task.Get("/", handler.GetAllTasks)
	task.Post("/", handler.CreateTask)
	task.Put("/:id", handler.UpdateTask)
	task.Delete("/:id", handler.DeleteTask)
}
