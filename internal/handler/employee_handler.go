package handler

import (
	"fmt"
	"net/http"

	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllEmployees(c *fiber.Ctx) error {
	var data []model.Employee
	err := config.DB.Find(&data).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Data not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   data,
	})
}

func GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Hello, World!!",
	})
}

func CreateEmployee(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Hello, World!",
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Hello, World!",
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Hello, World!",
	})
}
