package models

type Subject struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name    string `gorm:"type:text;not null"`
	ClassID uint   `gorm:"not null"`
	Class   Class  `gorm:"foreignKey:ClassID"`
}
