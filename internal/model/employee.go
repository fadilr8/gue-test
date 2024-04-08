package model

import (
	"time"
)

type Employee struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday *time.Time `json:"birthday" gorm:"type:date"`
	Domicile string    `json:"domicile"`
	Address  string    `gorm:"type:text" json:"address"`
}
