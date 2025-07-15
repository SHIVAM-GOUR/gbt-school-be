package models

type Class struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:text;not null"`
}
