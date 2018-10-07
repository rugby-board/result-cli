package match

// Retriever types
const (
	PlanetRugby = "planet-rugby"
	RugbyComAu  = "rugby-com-au"
)

const (
	// InternationalTests ...
	InternationalTests = 3
	// Premiership ...
	Premiership = 201
	// Top14 ...
	Top14 = 203
	// Pro14 ...
	Pro14 = 204
	// SuperRugby ...
	SuperRugby = 205
	// AngloWelshCup ...
	AngloWelshCup = 206
	// Mitre10Cup ...
	Mitre10Cup = 208
	// SixNations ...
	SixNations = 209
	// RugbyWorldCup ...
	RugbyWorldCup = 210
	// TheRugbyChampionship ...
	TheRugbyChampionship = 214
	// BritishAndIrishLions ...
	BritishAndIrishLions = 221
	// EuropeanChampionCup ...
	EuropeanChampionCup = 242
	// EuropeanChallengeCup ...
	EuropeanChallengeCup = 243
	// CurrieCupPremier ...
	CurrieCupPremier = 303
	// NationalRugbyChampionship ...
	NationalRugbyChampionship = 247
)

// Event of a tournament
type Event struct {
	Type string
	ID   int32
	Name string
}

var matchEvents = map[int32]string{
	InternationalTests:        "International Tests",
	Premiership:               "Premiership",
	Top14:                     "Top14",
	Pro14:                     "Pro14",
	SuperRugby:                "Super Rugby",
	AngloWelshCup:             "Anglo Welsh Cup",
	Mitre10Cup:                "Mitre10 Cup",
	SixNations:                "Six Nations",
	RugbyWorldCup:             "Rugby World Cup",
	TheRugbyChampionship:      "The Rugby Championship",
	BritishAndIrishLions:      "British & Irish Lions",
	EuropeanChampionCup:       "European Champion Cup",
	EuropeanChallengeCup:      "European Challenge Cup",
	CurrieCupPremier:          "Currie Cup Premier",
	NationalRugbyChampionship: "National Rugby Championship",
}

// ValidEvent ...
func ValidEvent(eventID int32) bool {
	_, ok := matchEvents[eventID]
	return ok
}

// GetEvent ...
func GetEvent(eventID int32) (*Event, bool) {
	name, ok := matchEvents[eventID]
	if ok {
		return &Event{ID: eventID, Name: name, Type: getEventType(eventID)}, true
	}
	return nil, false
}

func getEventType(eventID int32) string {
	if eventID == NationalRugbyChampionship {
		return RugbyComAu
	}
	return PlanetRugby
}

// ListEvents ...
func ListEvents() []Event {
	events := make([]Event, 0)
	for eventID, name := range matchEvents {
		events = append(events, Event{ID: eventID, Name: name, Type: getEventType(eventID)})
	}
	return events
}
