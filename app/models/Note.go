package models

import (
	_ "ZenZen_API/app"
	_ "ZenZen_API/framework/utils"
	_ "errors"
	_ "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
