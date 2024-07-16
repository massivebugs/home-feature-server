package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/route"
	"github.com/massivebugs/home-feature-server/internal/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatalf("ENV has not been specified, exiting")
	}

	port := os.Getenv("API_PORT")

	e := echo.New()
	e.Validator = &api.RequestValidator{}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: Switch CORS settings on local and production
	e.Use(middleware.CORS())

	route.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
