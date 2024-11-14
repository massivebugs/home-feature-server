package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/seeder"
)

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

	seeder := seeder.NewSeeder(
		db,
		queries.New(),
	)

	switch cfg.Environment {
	case rest.EnvironmentLocal:
		err = seeder.SeedForLocal(context.Background())
	case rest.EnvironmentProduction:
		err = seeder.SeedForProduction(context.Background())
	}
	if err != nil {
		log.Fatal(err)
	}
}
