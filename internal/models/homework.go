package models

import (
	"time"
)

type Homework struct {
	ID           uint       `gorm:"primaryKey;autoIncrement;not null"`
	Title        string     `gorm:"type:text;not null"`
	Description  string     `gorm:"type:text;not null"`
	ClassID      uint       `gorm:"not null"`
	SubjectID    uint       `gorm:"not null"`
	TeacherID    uint       `gorm:"not null"`
	AssignedDate *time.Time `gorm:"type:date"`
	CreatedAt    time.Time  `gorm:"type:timestamp"`
	UpdatedAt    time.Time  `gorm:"type:timestamp"`
	Class        Class      `gorm:"foreignKey:ClassID"`
	Subject      Subject    `gorm:"foreignKey:SubjectID"`
	Teacher      Teacher    `gorm:"foreignKey:TeacherID"`
}
