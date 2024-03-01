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
	return &User{
		Id:       uuid.New().String(),
		Username: username,
		LoginDetails: LoginDetails{
			Username:       username,
			HashedPassword: passwordHash,
		},
	}
}
