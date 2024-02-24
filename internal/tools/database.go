package tools

type LoginDetails struct {
	Username        string `json:"username"`
	AuthToken       string `json:"authToken"`
	FavouriteNumber int    `json:"favouriteNumber"`
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &gormDb{}

	var err error = database.SetupDatabase()
	if err != nil {
		return nil, err
	}

	return &database, nil
}
