package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      uint
	Name    string
	Age     uint
	Email   string
	Balance uint
}
