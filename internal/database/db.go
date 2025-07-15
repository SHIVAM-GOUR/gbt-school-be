package database

import (
	"fmt"

	"github.com/SHIVAM-GOUR/gbt-school-be/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
