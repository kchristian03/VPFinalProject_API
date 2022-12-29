package models

import (
	"ZenZen_API/framework/database"
)

func RegisteredModels() []Model {
	return []Model{
		{
			Name:   "User",
			Model:  User{},
			Seeder: nil,
			//&database.SeederDefinition{
			//Amount: 10,
			//Run: func(db *gorm.DB) error {
			//	password, err := app.Utilities.Crypto.HashPassword("password")
			//
			//	if err != nil {
			//		return err
			//	}
			//
			//	user := User{
			//		Name:     gofakeit.Name(),
			//		Email:    gofakeit.Email(),
			//		Username: gofakeit.Username(),
			//		Password: password,
			//	}
			//
			//	return db.Create(&user).Error
			//},
			//},
		},
		{
			Name:   "Personal Access Token",
			Model:  PersonalAccessToken{},
			Seeder: nil,
		},
		{
			Name:   "Music",
			Model:  Music{},
			Seeder: nil,
		},
		{
			Name:   "Note",
			Model:  Note{},
			Seeder: nil,
		},
		{
			Name:   "NoteCollection",
			Model:  NoteCollection{},
			Seeder: nil,
		},
		{
			Name:   "FavoriteMusic",
			Model:  FavoriteMusic{},
			Seeder: nil,
		},
	}
}

type Model struct {
	Name   string
	Model  interface{}
	Seeder *database.SeederDefinition
}
