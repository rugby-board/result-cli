package match

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
)

// ValidEvent ...
func ValidEvent(eventID int32) bool {
	if eventID == InternationalTests {
		return true
	} else if eventID == Premiership {
		return true
	} else if eventID == Top14 {
		return true
	} else if eventID == Pro14 {
		return true
	} else if eventID == SuperRugby {
		return true
	} else if eventID == AngloWelshCup {
		return true
	} else if eventID == Mitre10Cup {
		return true
	} else if eventID == SixNations {
		return true
	} else if eventID == RugbyWorldCup {
		return true
	} else if eventID == TheRugbyChampionship {
		return true
	} else if eventID == BritishAndIrishLions {
		return true
	} else if eventID == EuropeanChampionCup {
		return true
	} else if eventID == EuropeanChallengeCup {
		return true
	} else if eventID == CurrieCupPremier {
		return true
	}

	return false
}
