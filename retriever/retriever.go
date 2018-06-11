package retriever

import (
	"net/http"

	"github.com/rugby/result-cli/match"
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
func Init(r *Retriever) error {
	r.client = &http.Client{}
	return nil
}

// Retrieve results
func Retrieve(r *Retriever) error, match.Match {
	return nil
}
