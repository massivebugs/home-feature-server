package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/db/seeder"
	"github.com/massivebugs/home-feature-server/db/service/auth_repository"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

func main() {
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

	seeder.NewSeeder(
		db,
		cfg,
		auth_repository.New(),
		cashbunny_repository.NewCashbunnyRepository(),
	).Seed(context.Background())
}
