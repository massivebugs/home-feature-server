package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type Server struct {
	*rest.PingHandler
	*rest.RepeatHandler
	*rest.AuthHandler
	// *rest.SystemPreferencesHandler
}

var _ oapi.StrictServerInterface = (*Server)(nil)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Checking config...")
	cfg := rest.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("Creating database connection...")
	db, err := db.OpenMySQLDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	// log.Println("Fetching swagger specification...")
	// swagger, err := rest.GetSwagger()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Attaching middlewares...")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(rest.NewCSRFMiddleware(cfg))
	e.Use(rest.NewCORSMiddleware(cfg))
	e.Use(rest.NewJWTMiddleware(cfg))
	e.Use(rest.NewJWTRefreshMiddleware(cfg))
	e.HTTPErrorHandler = rest.NewHTTPErrorHandler(cfg)
	e.Validator = rest.NewRequestValidator()

	// TODO: Move JWT validation from JWTMiddleware to OapiRequestValidator
	// e.Use(echomiddleware.OapiRequestValidatorWithOptions(swagger, &echomiddleware.Options{
	// 	Options: openapi3filter.Options{
	// 		AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
	// 			// Return nil here because we'll be using authentication with Echo
	// 			return nil
	// 		},
	// 	},
	// }))

	log.Println("Registering server handlers...")
	querier := queries.New() //	query helpers generated from sqlc
	s := Server{
		PingHandler:   rest.NewPingHandler(cfg),
		RepeatHandler: rest.NewRepeatHandler(cfg),
		AuthHandler:   rest.NewAuthHandler(cfg, db, querier),
		// SystemPreferencesHandler: rest.NewSystemPreferencesHandler(cfg, db, querier),
		// CashbunnyHandler:             rest.NewCashbunnyHandler(cfg, db, querier),
	}

	oapi.RegisterHandlers(
		e,
		oapi.NewStrictHandler(
			s,
			[]oapi.StrictMiddlewareFunc{
				rest.RequestValidatorStrictHandlerFunc,
			}))

	e.Logger.Fatal(e.StartTLS(":"+cfg.APIPort, cfg.TLSCertificate, cfg.TLSKey))
}
