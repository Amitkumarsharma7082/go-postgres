package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primary Key;autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	// *string : When the field is optional, or might be null in JSON or DB
}

func MigrateBooks(db *gorm.DB) error {
	/*
		AutoMigrate is a Gorm method that automatically creates or updates the database table schema to match the structure of your Go struct (Books in this case).
	*/
	err := db.AutoMigrate(&Books{})
	return err
}
