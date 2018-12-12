package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rugby-board/go-rugby-dict/dict"
	"github.com/rugby-board/result-cli/cmd"
	"github.com/rugby-board/result-cli/match"
	"github.com/rugby-board/result-cli/retriever"
)

var (
	eventID    int
	daysBefore int
	round      int
	listEvents bool
	iterEvents bool
)

const defaultConfFile = "conf/conf.yaml"

var (
	planetRugbyRetriever = &retriever.PlanetRugbyRetriever{}
	rugbyComAuRetriever  = &retriever.RugbyComAuRetriever{}
	r                    retriever.Retriever
	d                    *dict.Dict
)

func main() {
	flag.IntVar(&eventID, "id", 0, "Event ID for Kratos")
	flag.IntVar(&daysBefore, "days", 7, "Days before")
	flag.IntVar(&round, "round", 1, "Round of game")
	flag.BoolVar(&listEvents, "list-events", false, "List events")
	flag.BoolVar(&iterEvents, "iter-events", false, "Iterate events")
	flag.Usage = usage
	flag.Parse()

	d = dict.NewDefaultDict()
	if d.Load() != nil {
		fmt.Println("load dict failed")
		return
	}

	realEventID := int32(eventID)
	dateStart, dateEnd := getDate(daysBefore)
	if listEvents {
		for _, event := range match.ListEvents() {
			fmt.Println(event)
		}
	} else if iterEvents {
		for _, event := range match.ListEvents() {
			retrieveResults(event, dateStart, dateEnd)
		}
	} else if match.ValidEvent(realEventID) {
		event, _ := match.GetEvent(realEventID)
		retrieveResults(*event, dateStart, dateEnd)
	} else {
		if realEventID > 0 {
			fmt.Println("Invalid event ID.")
			fmt.Println("Use -list-events to show all IDs.")
		} else {
			usage()
		}
	}
}

func retrieveResults(event match.Event, dateStart, dateEnd string) {
	color.Set(color.FgGreen)
	fmt.Println(event.Name)
	color.Unset()
	fmt.Printf("Event ID: %d, From %d days before: Fetching...\n\n", event.ID, daysBefore)
	if event.Type == match.RugbyComAu {
		r = rugbyComAuRetriever
	} else {
		r = planetRugbyRetriever
	}

	if r.Init(defaultConfFile) != nil {
		fmt.Println("Init retriever failed")
		return
	}
	var m []*match.Match
	if event.Type == match.RugbyComAu {
		tempDates := strings.Split(dateStart, "-")
		year, _ := strconv.Atoi(tempDates[0])
		m, _ = r.Retrieve(event.ID, fmt.Sprintf("%d", year+1), strconv.Itoa(round))
	} else {
		m, _ = r.Retrieve(event.ID, dateStart, dateEnd)
	}
	for _, item := range m {
		trans, err := d.Query(item.Team1Name)
		if err == nil {
			item.Team1Name = fmt.Sprintf("%s %s", trans, item.Team1Name)
		}
		trans, err = d.Query(item.Team2Name)
		if err == nil {
			item.Team2Name = fmt.Sprintf("%s %s", trans, item.Team2Name)
		}
	}
	if m != nil && len(m) != 0 {
		color.Set(color.FgGreen)
		fmt.Printf("Results:\n\tFirst game date: %s\n\n", m[0].GameDate)
		color.Unset()
		cmd.OutputMarkdownTable(m)
	} else {
		color.Set(color.FgYellow)
		fmt.Println("No match result found.")
		color.Unset()
	}
	fmt.Println("")
}

func getDate(daysBefore int) (string, string) {
	year, month, day := time.Now().Date()
	dateEnd := fmt.Sprintf("%d-%d-%d", year, month, day)
	year, month, day = time.Now().AddDate(0, 0, -daysBefore).Date()
	dateStart := fmt.Sprintf("%d-%d-%d", year, month, day)
	return dateStart, dateEnd
}

func usage() {
	fmt.Printf("result-cli v%s\n", version())
	fmt.Println("Rugby match result retriever")
	fmt.Println("")
	fmt.Println("result-cli -id=[EVENT_ID] -days=[DAYS] -round=[ROUND] (-list-events) (-iter-events)")
}

func version() string {
	return "1.6.0"
}
