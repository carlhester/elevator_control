package timer

import (
	"testing"
	"time"
)

func TestStep(t *testing.T) {
	// given
	var timer Timer
	timer.Cycle = 0
	timer.CycleSpeed = 1

	// when
	curTime := time.Now()
	count := timer.Step(curTime)

	//then
	want := 1
	got := count

	if want != got {
		t.Errorf("Didn't get expected. Want: %v, Got: %v", want, got)
	}
}
