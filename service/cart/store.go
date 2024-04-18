package cart

import "database/sql"

type Store struct {
	db *sql.DB
}

func (store *Store) NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
