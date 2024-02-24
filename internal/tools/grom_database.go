package tools

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gormDb struct {
	*gorm.DB
}

func (d *gormDb) GetUserLoginDetails(username string) *LoginDetails {
	var db *gorm.DB
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	var loginDetails LoginDetails
	db.First(&loginDetails, "username = ?", username)

	return &loginDetails
}

func (d *gormDb) SetupDatabase() error {
	var dsn string = "test.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate()
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}
