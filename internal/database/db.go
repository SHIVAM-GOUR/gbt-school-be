package database

import (
	"log"
	"os"
	"time"

	"github.com/SHIVAM-GOUR/gbt-school-be/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	logLevel := logger.Info

	dsn := os.Getenv("RAILWAY_DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(
		75,
	) // Set the maximum number of connections in the idle connection pool
	sqlDB.SetMaxOpenConns(
		300,
	) // Set the maximum number of open connections to the database
	sqlDB.SetConnMaxLifetime(
		30 * time.Minute,
	) // Set the maximum amount of time a connection may be reused
	sqlDB.SetConnMaxIdleTime(
		10 * time.Minute,
	) // Set the maximum amount of time a connection may be idle

	log.Println("[info]: successfully connected to database")
	return db, nil
}
