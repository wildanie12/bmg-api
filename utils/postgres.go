package utils

import (
	"fmt"
	_config "gin-bmg-restful/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresConnection creates database connection
// to defined postgres instance inside configuration
func NewPostgresConnection(config *_config.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta", 
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
