package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Winterson-Islary/jwt-golang.git/config"
	"github.com/Winterson-Islary/jwt-golang.git/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	db_config := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.SSLMode)
	db, err := db.NewSQLStorage(db_config)
	if err != nil {
		log.Fatal(err)
	}
	InitStorage(db)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	mgrt, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := mgrt.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := mgrt.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

func InitStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("DB:%s Connection Successful\n", config.Envs.DBPort)
}
