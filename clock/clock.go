package clock

import "time"

func (c *Clock) Step(now *time.Time) {
	delay, _ := time.ParseDuration("1s")
	for time.Since(*now) < delay {
		time.Sleep(100 * time.Millisecond)
		c.Cycle += 1
	}
}
