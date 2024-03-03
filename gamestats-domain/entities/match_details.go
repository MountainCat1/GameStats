package entities

type MatchType int

const (
	Normal MatchType = iota // iota is 0
)

type MatchDetails struct {
	ID          string    `json:"id"`
	DateStarted string    `json:"dateStarted"`
	DateEnded   string    `json:"dateEnded"`
	MatchType   MatchType `json:"matchType"`
}

func CreateMatch(matchType MatchType) *MatchDetails {
	return &MatchDetails{
		MatchType: matchType,
	}
}
