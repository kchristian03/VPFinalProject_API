package models

import (
	"gorm.io/gorm"
)

type NoteCollection struct {
	gorm.Model
	// Add your fields here
	Note   Note `gorm:"foreignKey:NoteID"`
	NoteID uint
	User   User `gorm:"foreignKey:UserID"`
	UserID uint
}
