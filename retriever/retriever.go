package retriever

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rugby-board/result-cli/conf"
	"github.com/rugby-board/result-cli/match"
)

// Retriever struct
type Retriever struct {
	client  *http.Client
	baseURL string
}

// NewRetriever returns a Retriever
func NewRetriever() *Retriever {
	return &Retriever{}
}

// Init initialize HTTP client
func (r *Retriever) Init(confPath string) error {
	r.client = &http.Client{}
	confBody, err := conf.GetConf(confPath)
	if err != nil {
		log.Fatalf("Read file failed: %#v", err)
	}
	r.baseURL = confBody.BaseURL
	return nil
}

// Retrieve results
func (r *Retriever) Retrieve(eventID int32, dateStart string, dateEnd string) ([]*match.Match, error) {
	m := make([]*match.Match, 0)
	url := fmt.Sprintf(r.baseURL, eventID, dateStart, dateEnd)
	resp, err := r.client.Get(url)
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&m)
	return m, err
}
