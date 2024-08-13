package storage

import (
	"database/sql"
	"errors"
)

var ErrorNotFound = errors.New("Record Not Found")

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
