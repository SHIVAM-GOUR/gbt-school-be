package models

import (
	"time"
)

type Announcement struct {
	ID            uint       `gorm:"primaryKey;autoIncrement;not null"`
	Title         string     `gorm:"type:text;not null"`
	Description   string     `gorm:"type:text;not null"`
	PublishedDate *time.Time `gorm:"type:date"`
	StartDate     *time.Time `gorm:"type:date"`
	EndDate       *time.Time `gorm:"type:date"`
	Audience      string     `gorm:"type:text"`
	AttachmentURL string     `gorm:"type:text"`
	IsActive      bool       `gorm:"type:boolean"`
	CreatedAt     time.Time  `gorm:"type:timestamp"`
	UpdatedAt     time.Time  `gorm:"type:timestamp"`
}
