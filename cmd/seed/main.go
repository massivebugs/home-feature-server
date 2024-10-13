package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/http"
	"github.com/massivebugs/home-feature-server/seeder"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Checking config...")
	cfg := http.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating database connection...")
	db, err := db.OpenMySQLDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)
	if err != nil {
		log.Fatal(err)
	}

	seeder := seeder.NewSeeder(
		db,
		queries.New(),
	)

	switch cfg.Environment {
	case http.EnvironmentLocal:
		err = seeder.SeedForLocal(context.Background())
	case http.EnvironmentProduction:
		err = seeder.SeedForProduction(context.Background())
	}
	if err != nil {
		log.Fatal(err)
	}
}
