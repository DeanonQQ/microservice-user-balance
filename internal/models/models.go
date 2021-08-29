package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      int64
	Name    string
	Age     int64
	Email   string
	Balance float64
}
