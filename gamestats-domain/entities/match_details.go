package entities

type MatchDetails struct {
	ID          string `json:"id"`
	DateStarted string `json:"dateStarted"`
	DateEnded   string `json:"dateEnded"`
	MatchType   string `json:"matchType"`
}
