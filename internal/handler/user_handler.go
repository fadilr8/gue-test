package handler

import (
	"net/http"
	"time"

	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/model"
	"github.com/fadilr8/gue-test/internal/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(model.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": false,
			"error":  err.Error(),
		})
	}

	exp := time.Hour * 24 * 30

	claims := jtoken.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(exp).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.SECRET))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"data": fiber.Map{
			"token": t,
		},
	})
}
