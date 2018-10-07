package retriever

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/rugby-board/result-cli/conf"
	"github.com/rugby-board/result-cli/match"
)

// RugbyComAuRetriever ...
type RugbyComAuRetriever struct {
	Client
}

// RugbyComAuResult ...
type RugbyComAuResult struct {
	Rounds []RugbyComAuRound `json:"rounds"`
}

// RugbyComAuRound ...
type RugbyComAuRound struct {
	Name      string            `json:"name"`
	RoundType string            `json:"round_type"`
	RoundName string            `json:"round_name"`
	Fixtures  []RugbyComAuMatch `json:"fixtures"`
}

// RugbyComAuMatch match struct of planet rugby
type RugbyComAuMatch struct {
	DateTime string           `json:"datetime"`
	CompID   string           `json:"comp_id"`
	SeasonID string           `json:"season_id"`
	MatchID  string           `json:"match_id"`
	Venue    string           `json:"venue"`
	Teams    []RugbyComAuTeam `json:"teams"`
}

// RugbyComAuTeam ...
type RugbyComAuTeam struct {
	Name  string `json:"name"`
	Score string `json:"score"`
	Crest string `json:"crest"`
}

// Init initialize HTTP client
func (r *RugbyComAuRetriever) Init(confPath string) error {
	r.client = &http.Client{}
	confBody, err := conf.GetConf(confPath)
	if err != nil {
		log.Fatalf("Read file failed: %#v", err)
	}
	r.baseURL = confBody.RugbyComAuURL
	return nil
}

// Retrieve results
func (r RugbyComAuRetriever) Retrieve(eventID int32, year string, round string) ([]*match.Match, error) {
	matches := RugbyComAuResult{}
	yearInt, _ := strconv.Atoi(year)
	url := fmt.Sprintf(r.baseURL, eventID, yearInt)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&matches)
	if len(matches.Rounds) == 0 {
		return nil, err
	}
	var roundResult RugbyComAuRound
	var isRoundFound bool
	for _, m := range matches.Rounds {
		if m.RoundName == round {
			roundResult = m
			isRoundFound = true
			break
		}
	}
	if !isRoundFound {
		return nil, err
	}
	results := make([]*match.Match, 0)
	for _, m := range roundResult.Fixtures {
		convertedMatch := r.ConvertMatchData(m)
		results = append(results, &convertedMatch)
	}
	return results, err
}

// ConvertMatchData ...
func (r RugbyComAuRetriever) ConvertMatchData(data interface{}) match.Match {
	m := data.(RugbyComAuMatch)
	team1Score, _ := strconv.Atoi(m.Teams[0].Score)
	team2Score, _ := strconv.Atoi(m.Teams[1].Score)
	t, _ := time.Parse(time.RFC3339, m.DateTime)
	convertedMatch := match.Match{
		Team1Name:  m.Teams[0].Name,
		Team1Score: int32(team1Score),
		Team2Name:  m.Teams[1].Name,
		Team2Score: int32(team2Score),
		Venue:      m.Venue,
		GameDate:   fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()),
		GameTime:   fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()),
	}
	return convertedMatch
}
