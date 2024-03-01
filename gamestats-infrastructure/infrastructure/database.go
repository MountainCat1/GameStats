package infrastructure

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"gorm.io/gorm"
	"time"
)

func NewDatabase(cfg Config) (*gorm.DB, error) {
	db, err := NewGormDb(cfg)
	if err != nil {
		return nil, err
	}

	err = MigrateGormDb(db, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewNoSqlDatabase(cfg Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cs, err := connstring.ParseAndValidate(cfg.Database.ConnectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(cs.Original)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Get a handle to the "your_database_name" database
	db := client.Database(cs.Database)

	return db, nil
}
