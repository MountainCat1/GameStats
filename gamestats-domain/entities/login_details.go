package entities

type LoginDetails struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	UserID         string `json:"userID"`
}
