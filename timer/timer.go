package timer

import "time"

func (t *Timer) Step(now *time.Time) {
	delay, _ := time.ParseDuration("1s")
	for time.Since(*now) < delay {
		time.Sleep(100 * time.Millisecond)
		t.Cycle += 1
	}
}
