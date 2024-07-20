package config

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type EnvironmentType string

const (
	EnvironmentLocal      EnvironmentType = "local"
	EnvironmentProduction EnvironmentType = "production"
)

type Config struct {
	Environment EnvironmentType
	APIPort     string
	JWTSecret   string

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
	c.Environment = EnvironmentType(os.Getenv("ENVIRONMENT"))
	c.APIPort = os.Getenv("API_PORT")
	c.JWTSecret = os.Getenv("JWT_SECRET")
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
			&c.JWTSecret,
			validation.Required,
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
