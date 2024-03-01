package infrastructure

import (
	"errors"
	"fmt"
	"gamestats-domain/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) AddUser(user *entities.User) (error, *entities.User) {
	tx := repo.db.Create(user)
	if tx.Error != nil {
		return fmt.Errorf("failed to add user: %w", tx.Error), nil
	}
	return nil, user
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) GetUserLoginDetails(username string) (*entities.LoginDetails, error) {
	db := repo.db

	var loginDetails entities.LoginDetails
	err := db.First(&loginDetails, "username = ?", username).Error

	if err == nil {
		return &loginDetails, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Record not found
		return nil, nil
	}

	// Other error occurred
	return nil, fmt.Errorf("failed to get user login details: %w", err)
}
