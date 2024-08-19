package storage

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorNotFound = errors.New("Record Not Found")

type Storage struct {
	//db *sql.DB
	db *gorm.DB
}

/*
func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
*/

func NewStore(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}
