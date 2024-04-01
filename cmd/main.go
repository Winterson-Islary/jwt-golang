package main

import (
	"fmt"
	"log"

	"github.com/Winterson-Islary/jwt-golang.git/cmd/api"
	"github.com/Winterson-Islary/jwt-golang.git/cmd/config"
	"github.com/Winterson-Islary/jwt-golang.git/cmd/db"
)

func main() {
	db_config := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.Envs.DBUser, config.Envs.DBName, config.Envs.DBPassword, config.Envs.SSLMode)
	db, err := db.NewSQLStorage(db_config)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
