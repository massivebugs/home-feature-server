package main

import (
	"log"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/route"
)

type CustomValidator struct{}

func (*CustomValidator) Validate(i interface{}) error {
	if v, ok := i.(validation.Validatable); ok {
		return v.Validate()
	}
	return nil
}

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
	e.Validator = &CustomValidator{}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: Switch CORS settings on local and production
	e.Use(middleware.CORS())

	route.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + port))
}
