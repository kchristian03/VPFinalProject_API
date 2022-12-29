package models

import (
	_ "ZenZen_API/app"
	_ "ZenZen_API/framework/utils"
	_ "errors"
	_ "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Duration int    `gorm:"not null"`
	Composer string `gorm:"not null"`
	Image    string `gorm:"not null"`
	Audio    string `gorm:"not null"`
}
