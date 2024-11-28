package main

import (
	"github.com/labstack/echo/v4"
	"github.com/priyanshoon/stemcms/internal/database"
	"net/http"
)

func main() {
	e := echo.New()
	db := database.MakeDatabase()
	db.CreateUsers("Priyanshu Prasad Gupta", "priyanshoon.pg@gmail.com", "1212")
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})
	e.Logger.Fatal(e.Start(":6969"))
}
