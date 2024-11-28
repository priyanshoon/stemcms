package database

import (
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

type User struct {
	*gorm.Model
	Name     string
	Email    string
	Password string
}
