package retriever

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rugby-board/result-cli/match"
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
func (r *Retriever) Retrieve(url string) ([]*match.Match, error) {
	m := make([]*match.Match, 0)
	resp, err := r.client.Get(url)
	if err != nil {
		return m, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&m)
	fmt.Println(m)
	return m, err
}
