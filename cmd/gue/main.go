package main

import (
	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDatabase()

	route.InitRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("404 - Not Found")
	})

	app.Listen(":3000")
}
