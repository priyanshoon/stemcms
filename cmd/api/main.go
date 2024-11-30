package main

import (
	"github.com/labstack/echo/v4"
	"github.com/priyanshoon/stemcms/internal/database"
	"github.com/priyanshoon/stemcms/internal/server"
)

func main() {
	e := echo.New()
	db := database.MakeDatabase()
	server.RegisterAndLoginRoutes(db, e)

	e.Logger.Fatal(e.Start(":6969"))
}
