package server

import (
	"net/http"

	// "log"

	"github.com/labstack/echo/v4"
	"github.com/priyanshoon/stemcms/internal/database"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := new(database.User)

		if err := c.Bind(&u); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		user := &database.User{
			Name:     u.Name,
			Username: u.Username,
			Email:    u.Email,
			Password: u.Password,
		}

		userExist := db.Where("username = ?", user.Username).First(&u)

		if userExist.Error == nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Username already exist",
			})
		}

		emailExist := db.Where("email = ?", user.Email).First(&u)

		if emailExist.Error == nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Email already exist",
			})
		}

		db.Create(user)

		return c.JSON(http.StatusCreated, user)
	}
}

func LoginUser(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {

		u := new(database.User)

		if err := c.Bind(&u); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Internal Server Error",
			})
		}

		user := &database.User{
			Username: u.Username,
			Email:    u.Email,
		}

		userExist := db.Where("username = ?", user.Username).First(&u)

		if userExist.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Username not found!",
			})
		}

		emailExist := db.Where("username = ?", user.Username).First(&u)

		if emailExist.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Email not found!",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login Successfull",
		})
	}
}
