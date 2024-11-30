package database

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name     string `gorm:"not null" json:"name" form:"name"`
	Username string `gorm:"unique; not null" json:"username" form:"username"`
	Email    string `gorm:"not null; unique" json:"email" form:"email"`
	Password string `gorm:"not null" json:"password" form:"password"`
}
