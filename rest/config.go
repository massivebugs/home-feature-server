package rest

import (
	"os"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt/v5"
)

type EnvironmentType string

const (
	EnvironmentLocal      EnvironmentType = "local"
	EnvironmentProduction EnvironmentType = "production"
)

type Config struct {
	Environment    EnvironmentType
	APIPort        string
	AllowedOrigins []string

	TLSCertificate string
	TLSKey         string

	// DateTime format used for all request/responses
	APIDateTimeFormat string

	AuthJWTCookieName       string
	AuthJWTSigningMethod    *jwt.SigningMethodHMAC
	AuthJWTSecret           string
	AuthJWTExpireSeconds    int
	RefreshJWTCookieName    string
	RefreshJWTSigningMethod *jwt.SigningMethodHMAC
	RefreshJWTSecret        string
	RefreshJWTExpireSeconds int

	DBHost     string
	DBPort     string
	DBDatabase string
	DBUser     string
	DBPassword string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	var err error

	c.Environment = EnvironmentType(os.Getenv("ENVIRONMENT"))
	c.APIPort = os.Getenv("API_PORT")
	c.AllowedOrigins = strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	c.TLSCertificate = "devcerts/" + os.Getenv("TLS_CERTIFICATE")
	c.TLSKey = "devcerts/" + os.Getenv("TLS_KEY")

	c.APIDateTimeFormat = os.Getenv("API_DATETIME_FORMAT")

	c.AuthJWTCookieName = os.Getenv("AUTH_JWT_COOKIE_NAME")
	authTokenSigningMethod := os.Getenv("AUTH_JWT_SIGNING_METHOD_HMAC")
	if authTokenSigningMethod != "" {
		switch authTokenSigningMethod {
		case "HS256":
			c.AuthJWTSigningMethod = jwt.SigningMethodHS256
		case "HS384":
			c.AuthJWTSigningMethod = jwt.SigningMethodHS384
		case "HS512":
			c.AuthJWTSigningMethod = jwt.SigningMethodHS512
		}
	}
	c.AuthJWTSecret = os.Getenv("AUTH_JWT_SECRET")
	c.AuthJWTExpireSeconds, err = strconv.Atoi(os.Getenv("AUTH_JWT_EXPIRE_SECONDS"))
	if err != nil {
		return err
	}

	c.RefreshJWTCookieName = c.AuthJWTCookieName + "_refresh"
	refreshTokenSigningMethod := os.Getenv("REFRESH_JWT_SIGNING_METHOD_HMAC")
	if refreshTokenSigningMethod != "" {
		switch refreshTokenSigningMethod {
		case "HS256":
			c.RefreshJWTSigningMethod = jwt.SigningMethodHS256
		case "HS384":
			c.RefreshJWTSigningMethod = jwt.SigningMethodHS384
		case "HS512":
			c.RefreshJWTSigningMethod = jwt.SigningMethodHS512
		}
	}
	c.RefreshJWTSecret = os.Getenv("REFRESH_JWT_SECRET")
	c.RefreshJWTExpireSeconds, err = strconv.Atoi(os.Getenv("REFRESH_JWT_EXPIRE_SECONDS"))
	if err != nil {
		return err
	}

	c.DBHost = os.Getenv("MYSQL_HOST")
	c.DBPort = os.Getenv("MYSQL_PORT")
	c.DBDatabase = os.Getenv("MYSQL_DATABASE")
	c.DBUser = os.Getenv("MYSQL_USER")
	c.DBPassword = os.Getenv("MYSQL_PASSWORD")

	return c.validate()
}

func (c *Config) validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(
			&c.Environment,
			validation.Required,
			validation.In(EnvironmentLocal, EnvironmentProduction),
		),
		validation.Field(
			&c.APIPort,
			validation.Required,
			is.Port,
		),
		validation.Field(
			&c.AllowedOrigins,
			validation.Required,
		),
		validation.Field(
			&c.TLSCertificate,
			validation.Required,
		),
		validation.Field(
			&c.TLSKey,
			validation.Required,
		),
		validation.Field(
			&c.APIDateTimeFormat,
			validation.Required,
			validation.By(IsValidDateTimeFormat),
		),
		validation.Field(
			&c.AuthJWTCookieName,
			validation.Required,
		),
		validation.Field(
			&c.AuthJWTSigningMethod,
			validation.Required,
			validation.In(
				*jwt.SigningMethodHS256,
				*jwt.SigningMethodHS384,
				*jwt.SigningMethodHS512,
			),
		),
		validation.Field(
			&c.AuthJWTSecret,
			validation.Required,
		),
		validation.Field(
			&c.AuthJWTExpireSeconds,
			validation.Required,
			validation.Min(1),
		),
		validation.Field(
			&c.RefreshJWTCookieName,
			validation.Required,
		),
		validation.Field(
			&c.RefreshJWTSigningMethod,
			validation.Required,
			validation.In(
				*jwt.SigningMethodHS256,
				*jwt.SigningMethodHS384,
				*jwt.SigningMethodHS512,
			),
		),
		validation.Field(
			&c.RefreshJWTSecret,
			validation.Required,
		),
		validation.Field(
			&c.RefreshJWTExpireSeconds,
			validation.Required,
			validation.Min(1),
		),
		validation.Field(
			&c.DBHost,
			validation.Required,
		),
		validation.Field(
			&c.DBPort,
			validation.Required,
			is.Port,
		),
		validation.Field(
			&c.DBDatabase,
			validation.Required,
		),
		validation.Field(
			&c.DBUser,
			validation.Required,
		),
		validation.Field(
			&c.DBPassword,
			validation.Required,
		),
	)
}
