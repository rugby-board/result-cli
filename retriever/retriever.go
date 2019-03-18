package retriever

import (
	"net/http"

	"github.com/rugby-board/rugby-result/match"
)

// Retriever interface
type Retriever interface {
	Init(string) error
	Retrieve(int32, string, string) ([]*match.Match, error)
	ConvertMatchData(interface{}) match.Match
}

// Client of a retriever
type Client struct {
	client  *http.Client
	baseURL string
}
