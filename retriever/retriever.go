package retriever

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rugby-board/result-cli/match"
)

const (
	// BaseURL ...
	BaseURL = "http://kratos.365.co.za:9001/getresultsbycompidanddaterange/%d/%s/%s"
)

// Retriever struct
type Retriever struct {
	client *http.Client
}

// NewRetriever returns a Retriever
func NewRetriever() *Retriever {
	return &Retriever{}
}

// Init initialize HTTP client
func (r *Retriever) Init() error {
	r.client = &http.Client{}
	return nil
}

// Retrieve results
func (r *Retriever) Retrieve(eventID int32, dateStart string, dateEnd string) ([]*match.Match, error) {
	m := make([]*match.Match, 0)
	url := fmt.Sprintf(BaseURL, eventID, dateStart, dateEnd)
	resp, err := r.client.Get(url)
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&m)
	return m, err
}
