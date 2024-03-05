package value_types

type PlayerStats struct {
	PlayerId string `json:"playerId"`
	Kills    int    `json:"kills"`
	Deaths   int    `json:"deaths"`
	Score    int    `json:"score"`
}
