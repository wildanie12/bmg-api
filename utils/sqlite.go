package utils

import (
	_config "gin-bmg-restful/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewSQLiteConnection creates test database connection
// predefined test database credential using app config
func NewSQLiteConnection(config *_config.AppConfig) *gorm.DB {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}