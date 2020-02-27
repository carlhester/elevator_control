package clock

import (
	"testing"
	"time"
)

// TestStep tests that calling step takes at least 1 second, and adds at least 9 cycles
func TestStep(t *testing.T) {
	// given
	var clock Clock
	clock.Cycle = 0

	now := time.Now()

	// when
	clock.Step(&now)

	//then
	want, _ := time.ParseDuration("1s")
	got := time.Since(now)

	if want >= got {
		t.Errorf("Didn't get expected. Want: %v, Got: %v", want, got)
	}

	wantCycle := 9
	gotCycle := clock.Cycle

	if wantCycle >= gotCycle {
		t.Errorf("Didn't get expected. Want: %v, Got: %v", wantCycle, gotCycle)
	}
}
