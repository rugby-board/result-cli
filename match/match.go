package match

// Match struct
type Match struct {
	// Team info
	Team1Name  string `json:"team_name_1"`
	Team1Score int32  `json:"team_score_1"`
	Team2Name  string `json:"team_name_2"`
	Team2Score int32  `json:"team_score_2"`
	// Event info
	CompetitionName string `json:"competition_name"`
	PoolName        string `json:"pool_name"`
	Round           int32  `json:"round"`
	Referee         string `json:"referee"`
	Venue           string `json:"venue"`
	GameDate        string `json:"game_date"`
	GameTime        string `json:"game_time"`
}
