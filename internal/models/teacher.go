package models

import (
	"time"
)

type Teacher struct {
	ID          uint       `gorm:"primaryKey;autoIncrement;not null"`
	FirstName   string     `gorm:"type:text;not null"`
	LastName    string     `gorm:"type:text"`
	Email       string     `gorm:"type:text"`
	PhoneNumber string     `gorm:"type:text"`
	DateOfBirth *time.Time `gorm:"type:date"`
	Gender      string     `gorm:"type:text"`
	HireDate    *time.Time `gorm:"type:date"`
	Address     string     `gorm:"type:text"`
}
