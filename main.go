package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/go-template-api/config"
	"github.com/lucas-simao/go-template-api/internal/repositories"
)

func main() {
	env := config.NewConfig()

	e := echo.New()

	db := repositories.NewDB(env.DataSourceUrl)

	// Initialize controllers here for example: controllers.NewUserController(e, db)

	e.Logger.Fatal(e.Start(":9000"))
}
