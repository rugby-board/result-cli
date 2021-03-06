package match

import (
	"testing"
)

func TestValidEvent(t *testing.T) {
	validEvents := []int32{3, 201, 203, 204, 205, 206, 208, 209, 210, 214, 221, 242, 243, 303}
	for _, eventID := range validEvents {
		if ValidEvent(eventID) != true {
			t.Error("Test event failed on: ", eventID)
		}
	}
}

func TestInvalidEvent(t *testing.T) {
	if ValidEvent(4) != false {
		t.Error("Test event 4 failed")
	}
}

func TestListEvents(t *testing.T) {
	events := ListEvents()
	if len(events) != 15 {
		t.Error("Test events failed")
	}
}

func TestGetEvent(t *testing.T) {
	_, ok := GetEvent(777)
	if ok {
		t.Error("No event 777")
	}
	event, _ := GetEvent(247)
	if event.Type != RugbyComAu {
		t.Error("Event type error")
	}
}
