package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"test": {
		Username:        "test",
		AuthToken:       "123",
		FavouriteNumber: 5,
	},
	"test2": {
		Username:        "test2",
		AuthToken:       "1234",
		FavouriteNumber: 2,
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
