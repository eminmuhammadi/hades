package db

import (
	"fmt"
	"os"
	"time"

	_time "github.com/eminmuhammadi/hades/time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DSN = fmt.Sprintf(
	"host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=%v",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_SSLMODE"),
	os.Getenv("TIMEZONE"),
)

var Config = &gorm.Config{
	// Logger
	Logger: logger.Default.LogMode(logger.LogLevel(logger.Info)),
	// GORM perform write (create/update/delete) operations run
	// inside a transaction to ensure data consistency, which is
	// bad for performance, you can disable it during initialization
	SkipDefaultTransaction: true,
	// Creates a prepared statement when executing any
	// SQL and caches them to speed up future calls
	PrepareStmt: true,
	// QueryFields executes the SQL query with all fields of the table
	QueryFields: true,
	// NowFunc the function to be used when creating a new timestamp
	NowFunc: func() time.Time {
		return _time.Now()
	},
}

// Initialize the database
func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN), Config)

	return db, err
}

// Open the database connection
func Open() *gorm.DB {
	db, err := Initialize()

	// if there is an error, panic
	if err != nil {
		panic(err)
	}

	return db
}

// Close the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	return sqlDB.Close()
}
