package model

import (
	_time "github.com/eminmuhammadi/hades/time"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Log struct {
	DB_REQUIRED_COLS

	Data string `gorm:"type:text" json:"data"`
}

// Before create
func (model *Log) BeforeCreate(tx *gorm.DB) error {
	currentTime := _time.Now()

	// Set UUID for ID
	model.ID = uuid.NewV4()

	// Set timestamp for base
	model.CreatedAt = currentTime
	model.UpdatedAt = currentTime

	return nil
}

// Before update
func (model *Log) BeforeUpdate(tx *gorm.DB) error {
	currentTime := _time.Now()

	// Update timestamp for updated_at
	model.UpdatedAt = currentTime

	return nil
}

// Before delete
func (model *Log) BeforeDelete(tx *gorm.DB) error {
	currentTime := _time.Now()

	// Set timestamp for deleted_at
	model.DeletedAt.Time = currentTime

	return nil
}

// Table name
func (Log) TableName() string {
	return "logs"
}
