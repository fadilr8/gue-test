package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllEmployees(c *fiber.Ctx) error {
	var data []model.Employee

	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("limit", 10)
	offset := (page - 1) * pageSize

	err := config.DB.Limit(pageSize).Offset(offset).Find(&data).Error

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
	var data model.Employee

	err := config.DB.First(&data, "id =?", id).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Data not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"data":   data,
	})
}

func CreateEmployee(c *fiber.Ctx) error {
	var employee model.Employee
	var employeeRequest model.CreateEmployeeRequest

	if err := c.BodyParser(&employeeRequest); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	birthdayString := employeeRequest.Birthday
	if birthdayString != "" {
		var err error
		employee.Birthday, err = time.Parse("2006-01-02", birthdayString)
		if err != nil {
			return c.Status(400).SendString("Invalid birthday format (YYYY-MM-DD expected)")
		}
	}

	employee.Name = employeeRequest.Name
	employee.Email = employeeRequest.Email
	employee.Domicile = employeeRequest.Domicile
	employee.Address = employeeRequest.Address

	fmt.Println(birthdayString)
	if err := config.DB.Create(&employee).Error; err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "Employee created successfully!",
		"data":    employee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee model.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  false,
				"message": "Employee not found",
			})
		}
		return err
	}

	var employeeRequest model.CreateEmployeeRequest
	if err := c.BodyParser(&employeeRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if employeeRequest.Name != "" {
		employee.Name = employeeRequest.Name
	}
	if employeeRequest.Email != "" {
		employee.Email = employeeRequest.Email
	}
	if employeeRequest.Domicile != "" {
		employee.Domicile = employeeRequest.Domicile
	}
	if employeeRequest.Address != "" {
		employee.Address = employeeRequest.Address
	}

	birthdayString := employeeRequest.Birthday
	if birthdayString != "" {
		birthday, err := time.Parse("2006-01-02", birthdayString)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid birthday format (YYYY-MM-DD expected)")
		}
		employee.Birthday = birthday
	}

	if err := config.DB.Save(&employee).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Employee updated successfully",
		"data":    employee,
	})
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&model.Employee{}, id).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Employee deleted successfully",
	})
}
