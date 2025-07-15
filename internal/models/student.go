package models

import (
	"time"
)

type Student struct {
	StudentID      uint       `gorm:"primaryKey;autoIncrement;not null"`
	RollNumber     string     `gorm:"type:text"`
	FirstName      string     `gorm:"type:text;not null"`
	LastName       string     `gorm:"type:text"`
	DateOfBirth    *time.Time `gorm:"type:date"`
	Gender         string     `gorm:"type:text"`
	Email          string     `gorm:"type:text"`
	PhoneNumber    string     `gorm:"type:text"`
	Address        string     `gorm:"type:text"`
	EnrollmentDate *time.Time `gorm:"type:date"`
	ClassID        uint       `gorm:"not null"`
	Class          Class      `gorm:"foreignKey:ClassID"`
	Status         string     `gorm:"type:text"`
	CreatedAt      time.Time  `gorm:"type:timestamp"`
	UpdatedAt      time.Time  `gorm:"type:timestamp"`
}
