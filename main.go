package main

import "fmt"
import "elevator_control/elevator"
import "elevator_control/clock"
import "time"

func main() {

	var e1 elevator.Elevator
	e1.ID = 1
	e1.CurrentFloor = 1
	e1.DestFloor = 1
	e1.DoorOpen = false

	var clock clock.Clock
	clock.Cycle = 0

	for {
		// start clock
		clockTimer := time.Now()

		//refresh
		fmt.Printf("[%v] ID: %d, Current: %d, Dest: %d, DoorOpen: %v\n", clock.Cycle, e1.ID, e1.CurrentFloor, e1.DestFloor, e1.DoorOpen)

		// force sync to clock
		clock.Step(&clockTimer)
	}

}
