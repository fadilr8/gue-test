package route

import (
	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/handler"
	"github.com/fadilr8/gue-test/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App) {
	api := r.Group("/api")
	jwt := middleware.NewAuthMiddleware(config.SECRET)

	api.Post("/login", handler.Login)

	api.Get("/employees", jwt, handler.GetAllEmployees)
	api.Get("/employees/:id", jwt, handler.GetEmployee)
	api.Post("/employees", jwt, handler.CreateEmployee)
	api.Patch("/employees/:id", jwt, handler.UpdateEmployee)
	api.Delete("/employees/:id", jwt, handler.DeleteEmployee)
}
