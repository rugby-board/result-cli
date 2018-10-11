package main

import (
	"testing"

	"github.com/rugby-board/go-rugby-dict/dict"
	"github.com/rugby-board/result-cli/match"
)

func TestVersion(t *testing.T) {
	if len(version()) < 5 {
		t.Error("Illegal version")
	}
	usage()
}

func TestFetchPlanetRugby(t *testing.T) {
	daysBefore = 7
	round = 6
	d = dict.NewDefaultDict()
	if d.Load() != nil {
		t.Error("load dict failed")
		return
	}
	var event = match.Event{
		ID:   match.SuperRugby,
		Type: match.PlanetRugby,
	}
	s, e := getDate(daysBefore)
	retrieveResults(event, s, e)
	event = match.Event{
		ID:   match.NationalRugbyChampionship,
		Type: match.RugbyComAu,
	}
	s, e = getDate(daysBefore)
	retrieveResults(event, s, e)
}
