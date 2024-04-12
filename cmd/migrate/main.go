package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Winterson-Islary/jwt-golang.git/config"
	"github.com/Winterson-Islary/jwt-golang.git/db"
)

func main() {
	db_config := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName, config.Envs.SSLMode)
	db, err := db.NewSQLStorage(db_config)
	if err != nil {
		log.Fatal(err)
	}
	InitStorage(db)
}

func InitStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Printf("DB:%s Connection Successful\n", config.Envs.DBPort)
}
