package main

import "fmt"
import "elevator_control/elevator"
import "elevator_control/timer"
import "time"

func main() {
	start := time.Now()

	var e1 elevator.Elevator
	e1.ID = 1
	e1.CurrentFloor = 1
	e1.DestFloor = 1
	e1.DoorOpen = false

	var timer timer.Timer
	timer.Cycle = 0

	for {
		// update
		timer.Step(&start)

		// refresh
		fmt.Println(timer.Cycle, e1)
	}

}
