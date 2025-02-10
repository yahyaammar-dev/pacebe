package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteStorage() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("pace.db"), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}
	return db, nil
}
