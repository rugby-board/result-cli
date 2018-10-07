package retriever

import (
	"testing"
)

const defaultConfFile = "conf/conf.yaml"

func TestPlanetRugbyRetriever(t *testing.T) {
	r := PlanetRugbyRetriever{}
	if r.Init(defaultConfFile) != nil {
		t.Error("Init retriever failed")
	}
	m, err := r.Retrieve(205, "2017-06-11", "2018-06-11")
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}

func TestRugbyComAuRetriever(t *testing.T) {
	r := RugbyComAuRetriever{}
	if r.Init(defaultConfFile) != nil {
		t.Error("Init retriever failed")
	}
	m, err := r.Retrieve(247, "", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}
