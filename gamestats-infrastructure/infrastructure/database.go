package infrastructure

import "gorm.io/gorm"

type Database struct {
	gorm.DB
}

func NewDatabase(cfg Config) (*Database, error) {
	var gormDb = &gormDb{}

	var err error = gormDb.SetupDatabase(cfg)
	if err != nil {
		return nil, err
	}

	var database = Database{
		DB: *gormDb.DB, // TODO should it be a pointer?
	}

	return &database, nil
}
