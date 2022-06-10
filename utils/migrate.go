package utils

import (
	_domain "gin-bmg-restful/entities/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&_domain.Member{},
	)
}