package publish

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// PublishURL url of publishing
const PublishURL = "https://rugby-board.herokuapp.com/api/v1/news?token=%s"

// Publisher definition
type Publisher struct {
	client *http.Client
	url    string
	token  string
}

// NewPublisher creates a publisher
func NewPublisher() *Publisher {
	return &Publisher{}
}

// Init ...
func (p *Publisher) Init() {
	p.url = PublishURL
	p.client = &http.Client{}
	p.token = os.Getenv("API_TOKEN")
}

// Publish ...
func (p Publisher) Publish(title, content string, eventID int32) error {
	form := url.Values{}
	form.Add("title", title)
	form.Add("content", content)
	form.Add("channel", "1")
	form.Add("event", fmt.Sprintf("%d", eventID))
	_, err := p.client.PostForm(fmt.Sprintf(p.url, p.token), form)
	return err
}
