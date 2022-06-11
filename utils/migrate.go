package utils

import (
	_domain "gin-bmg-restful/entities/domain"

	"gorm.io/gorm"
)

// Migrate specified entity to database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&_domain.User{},
	)
}