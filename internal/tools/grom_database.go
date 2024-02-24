package tools

import (
	"GameStats/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gormDb struct {
	DB *gorm.DB
}

func (d *gormDb) SetupDatabase(cfg Config) error {
	var dsn string = cfg.Database.DBName + ".db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&entities.LoginDetails{})
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}
