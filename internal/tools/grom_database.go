package tools

import (
	"GameStats/internal/entities"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gormDb struct {
	DB *gorm.DB
}

var ErrInvalidDatabaseProvider = errors.New("invalid database provider")

func (d *gormDb) SetupDatabase(cfg Config) error {
	db, err := NewGormDb(cfg)
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

func NewGormDb(cfg Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch cfg.Database.Provider {
	case "sqlite":
		var connectionString = cfg.Database.ConnectionString
		db, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	case "mysql":
		var connectionString = cfg.Database.ConnectionString
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	default:
		return nil, ErrInvalidDatabaseProvider
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
