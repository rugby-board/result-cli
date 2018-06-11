package match

// Match struct
type Match struct {
	// Team info
	Team1Name  string
	Team1Score int32
	Team2Name  string
	Team2Score int32
	// Event info
	CompetitionName string
	PoolName        string
	Round           int32
	Referee         string
	Venue           string
	GameDate        string
	GameTime        string
}
