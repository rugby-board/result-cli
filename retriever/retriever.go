package retriever

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rugby-board/result-cli/match"
	yaml "gopkg.in/yaml.v2"
)

// Retriever struct
type Retriever struct {
	client  *http.Client
	baseURL string
}

type conf struct {
	BaseURL string `yaml:"base_url"`
}

// NewRetriever returns a Retriever
func NewRetriever() *Retriever {
	return &Retriever{}
}

// Init initialize HTTP client
func (r *Retriever) Init(confPath string) error {
	r.client = &http.Client{}
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("Read file failed: %#v", err)
		return err
	}
	c := &conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal failed: %#v", err)
		return err
	}
	r.baseURL = c.BaseURL
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
