package db

import (
	model "github.com/eminmuhammadi/hades/model"
	"gorm.io/gorm"
)

// Migrate the database
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Log{},
	)
}

// Auto migrates the database
func AutoMigrate(db *gorm.DB) error {
	return Migrate(db)
}
