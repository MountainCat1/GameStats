package tools

import (
	"GameStats.Domain/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *Database
}

func NewUserRepo(db *Database) *UserRepo {
	return &UserRepo{db}
}

func (repo *UserRepo) GetUserLoginDetails(username string) (*entities.LoginDetails, error) {
	db := repo.db.DB

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
