package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/rest"
)

type apiMiddlewares struct {
	CSRF       echo.MiddlewareFunc
	CORS       echo.MiddlewareFunc
	JWT        echo.MiddlewareFunc
	JWTRefresh echo.MiddlewareFunc
}

func main() {
	// TODO: Only in local?
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Checking config...")
	cfg := rest.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating database connection...")
	db, err := db.OpenMySQLDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &rest.RequestValidator{}

	fmt.Println("Attaching middlewares...")

	apiMiddlewares := apiMiddlewares{
		CSRF:       rest.NewCSRFMiddleware(cfg),
		CORS:       rest.NewCORSMiddleware(cfg),
		JWT:        rest.NewJWTMiddleware(cfg),
		JWTRefresh: rest.NewJWTRefreshMiddleware(cfg),
	}

	// Globally applied middleware
	// Route based middlewares can be applied at RegisterRoutes()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(apiMiddlewares.CORS)

	fmt.Println("Registering routes...")
	registerRoutes(e, cfg, apiMiddlewares, db)

	e.Logger.Fatal(e.StartTLS(":"+cfg.APIPort, cfg.TLSCertificate, cfg.TLSKey))
}
