package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string     `gorm:"primary_key;type:uuid;" json:"ID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (b *Base) BeforeCreate(scope *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return nil
}
