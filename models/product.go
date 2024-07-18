package models

type Product struct {
	Base
	UserID      string `gorm:"type:uuid;not null" json:"user_id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:1024;nor null" json:"description"`
	File        string `gorm:"size:255;not null" json:"file"`
	Version     string `gorm:"size:10;not null;" json:"version"`
}
