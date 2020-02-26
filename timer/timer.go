package timer

import "time"

func (t *Timer) Step(now *time.Time) int {
	target := now.Add(time.Second * 1)
	for now.Before(target) {
		time.Sleep(1 * time.Millisecond)
		*now = time.Now()
	}

	t.Cycle += 1
	return t.Cycle
}
