package entities

import (
	"github.com/google/uuid"
)

type User struct {
	Id           string       `json:"id"`
	Username     string       `json:"username"`
	LoginDetails LoginDetails `json:"loginDetails" gorm:"foreignKey:UserID"`
}

func CreateUser(username, passwordHash string) *User {
	id := uuid.New().String()

	return &User{
		Id:       id,
		Username: username,
		LoginDetails: LoginDetails{
			UserID:         id,
			Username:       username,
			HashedPassword: passwordHash,
		},
	}
}
