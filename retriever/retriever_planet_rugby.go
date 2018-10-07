package retriever

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rugby-board/result-cli/conf"
	"github.com/rugby-board/result-cli/match"
)

// PlanetRugbyRetriever ...
type PlanetRugbyRetriever struct {
	Client
}

// PlanetRugbyMatch match struct of planet rugby
type PlanetRugbyMatch struct {
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

// Init initialize HTTP client
func (r *PlanetRugbyRetriever) Init(confPath string) error {
	r.client = &http.Client{}
	confBody, err := conf.GetConf(confPath)
	if err != nil {
		log.Fatalf("Read file failed: %#v", err)
	}
	r.baseURL = confBody.PlanetRugbyURL
	return nil
}

// Retrieve results
func (r PlanetRugbyRetriever) Retrieve(eventID int32, dateStart string, dateEnd string) ([]*match.Match, error) {
	matches := make([]*PlanetRugbyMatch, 0)
	url := fmt.Sprintf(r.baseURL, eventID, dateStart, dateEnd)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&matches)
	results := make([]*match.Match, 0)
	for _, m := range matches {
		convertedMatch := r.ConvertMatchData(m)
		results = append(results, &convertedMatch)
	}
	return results, err
}

// ConvertMatchData ...
func (r PlanetRugbyRetriever) ConvertMatchData(data interface{}) match.Match {
	m := data.(*PlanetRugbyMatch)
	convertedMatch := match.Match{
		Team1Name:       m.Team1Name,
		Team1Score:      m.Team1Score,
		Team2Name:       m.Team2Name,
		Team2Score:      m.Team2Score,
		CompetitionName: m.CompetitionName,
		PoolName:        m.PoolName,
		Round:           m.Round,
		Referee:         m.Referee,
		Venue:           m.Venue,
		GameDate:        m.GameDate,
		GameTime:        m.GameTime,
	}
	return convertedMatch
}
