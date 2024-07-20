package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/api/route"
	"github.com/massivebugs/home-feature-server/internal/api"
)

func main() {
	fmt.Println("Checking config...")
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating database connection...")
	db, err := config.CreateDatabaseConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &api.RequestValidator{}

	fmt.Println("Attaching middlewares...")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// TODO: Switch CORS settings on local and production
	e.Use(middleware.CORS())

	fmt.Println("Registering routes...")
	route.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":" + cfg.APIPort))
}
