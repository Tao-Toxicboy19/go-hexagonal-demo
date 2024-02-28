package repositorys

import (
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

// new database
func NewDB(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}
