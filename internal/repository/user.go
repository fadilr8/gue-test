package repository

import (
	"errors"
	"fmt"

	"github.com/fadilr8/gue-test/internal/config"
	"github.com/fadilr8/gue-test/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func FindByCredentials(email, password string) (*model.CleanedUser, error) {
	var user model.User

	err := config.DB.First(&user, "email =?", email).Error

	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	clUser := model.CleanedUser{ID: user.ID, Name: user.Name, Email: user.Email}

	return &clUser, nil
}
