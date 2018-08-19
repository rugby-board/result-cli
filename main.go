package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/rugby-board/go-rugby-dict/dict"
	"github.com/rugby-board/result-cli/cmd"
	"github.com/rugby-board/result-cli/match"
	"github.com/rugby-board/result-cli/retriever"
)

var (
	eventID    int
	daysBefore int
	listEvents bool
)

const defaultConfFile = "conf/conf.yaml"

func main() {
	flag.IntVar(&eventID, "id", 0, "Event ID for Kratos")
	flag.IntVar(&daysBefore, "days", 7, "Days before")
	flag.BoolVar(&listEvents, "list-events", false, "List events")
	flag.Usage = usage
	flag.Parse()

	r := retriever.NewRetriever()
	r.Init(defaultConfFile)
	realEventID := int32(eventID)
	dateStart, dateEnd := getDate(daysBefore)
	if listEvents {
		for _, event := range match.ListEvents() {
			fmt.Println(event)
		}
	} else if match.ValidEvent(realEventID) {
		fmt.Printf("Event ID: %d, From %d days before:\n\n", realEventID, daysBefore)
		fmt.Printf("Fetching...\n\n")
		m, _ := r.Retrieve(realEventID, dateStart, dateEnd)
		d := dict.NewDefaultDict()
		err := d.Load()
		if err != nil {
			fmt.Println("Load dict failed")
		}
		var trans string
		for _, item := range m {
			trans, err = d.Query(item.Team1Name)
			if err == nil {
				item.Team1Name = fmt.Sprintf("%s %s", trans, item.Team1Name)
			}
			trans, err = d.Query(item.Team2Name)
			if err == nil {
				item.Team2Name = fmt.Sprintf("%s %s", trans, item.Team2Name)
			}
		}
		cmd.OutputMarkdownTable(m)
	} else {
		if realEventID > 0 {
			fmt.Println("Invalid event ID.")
			fmt.Println("Use -list-events to show all IDs.")
		} else {
			usage()
		}
	}
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
	fmt.Println("result-cli -id=[EVENT_ID] -days=[DAYS] (-list-events)")
}

func version() string {
	return "1.3.0"
}
