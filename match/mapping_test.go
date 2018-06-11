package match

import (
	"testing"
)

func TestValidEvent(t *testing.T) {
	if ValidEvent(3) != true {
		t.Error("Test event 3 failed")
	}
	if ValidEvent(4) != false {
		t.Error("Test event 4 failed")
	}
}
