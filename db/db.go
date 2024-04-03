package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewSQLStorage(db_config string) (*sql.DB, error) {
	db, err := sql.Open("postgres", db_config)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
