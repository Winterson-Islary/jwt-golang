package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Winterson-Islary/jwt-golang.git/cmd/api"
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

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DB:%s Connection Successful!\n", config.Envs.DBPort)
}
