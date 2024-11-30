package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAndLoginRoutes(db *gorm.DB, e *echo.Echo) {
	e.POST("/register", RegisterUser(db))
	e.POST("/login", LoginUser(db))
}
