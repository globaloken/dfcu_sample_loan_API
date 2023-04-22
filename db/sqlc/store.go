package db

import (
	"database/sql"
)

// store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
