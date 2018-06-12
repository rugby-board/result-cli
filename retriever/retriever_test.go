package retriever

import (
	"testing"
)

func TestRetriever(t *testing.T) {
	r := NewRetriever()
	if r.Init() != nil {
		t.Error("Init retriever failed")
	}
	m, err := r.Retrieve(205, "2017-06-11", "2018-06-11")
	if err != nil {
		t.Error(err)
	}
	t.Log(m)
}
