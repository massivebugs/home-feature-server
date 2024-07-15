package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	// Parse command line flags
	var isRollback bool
	flag.BoolVar(&isRollback, "rollback", false, "Rolls back a step if this is provided")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatalf("ENV has not been specified, exiting")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Reference - https://github.com/golang-migrate/migrate/tree/master/database/mysql
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+wd+"/db/migration",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if isRollback {
		err = m.Steps(-1)
	} else {
		err = m.Up()
	}

	if err != nil {
		log.Fatal(err)
	}
}
