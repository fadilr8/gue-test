package route

import (
	"github.com/fadilr8/gue-test/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App) {
	api := r.Group("/api")

	api.Post("/login", handler.Login)

	api.Get("/employees", handler.GetAllEmployees)
	api.Get("/employees/:id", handler.GetEmployee)
	api.Post("/employees", handler.CreateEmployee)
	api.Patch("/employees/:id", handler.UpdateEmployee)
	api.Delete("/employees/:id", handler.DeleteEmployee)
}
