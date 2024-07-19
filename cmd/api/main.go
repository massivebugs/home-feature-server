package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/route"
	"github.com/massivebugs/home-feature-server/internal/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := config.CreateDatabaseConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &api.RequestValidator{}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: Switch CORS settings on local and production
	e.Use(middleware.CORS())

	route.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + cfg.APIPort))
}
