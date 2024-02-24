package entities

type LoginDetails struct {
	Username        string `json:"username"`
	AuthToken       string `json:"authToken"`
	FavouriteNumber int    `json:"favouriteNumber"`
}
