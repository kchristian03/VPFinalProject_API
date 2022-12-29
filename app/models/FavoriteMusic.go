package models

import (
	_ "ZenZen_API/app"
	_ "ZenZen_API/framework/utils"
	_ "errors"
	_ "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FavoriteMusic struct {
	gorm.Model
	// Add your fields here
	Music   Music `gorm:"foreignKey:MusicID"`
	MusicID uint
	User    User `gorm:"foreignKey:UserID"`
	UserID  uint
}
