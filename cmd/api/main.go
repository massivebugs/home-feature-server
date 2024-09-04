package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/api/config"
	api_middleware "github.com/massivebugs/home-feature-server/api/middleware"
	"github.com/massivebugs/home-feature-server/api/route"
	"github.com/massivebugs/home-feature-server/internal/api"
)

func main() {
	// TODO: Only in local?
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	apiMiddleware := api_middleware.NewAPIMiddleware(cfg)

	// Globally applied middleware
	// Route based middlewares can be applied at RegisterRoutes()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(apiMiddleware.CORS)

	fmt.Println("Registering routes...")
	route.RegisterRoutes(e, cfg, apiMiddleware, db)

	e.Logger.Fatal(e.StartTLS(":"+cfg.APIPort, cfg.TLSCertificate, cfg.TLSKey))
}
