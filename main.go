package main

import (
	"elevator_control/clock"
	"elevator_control/controls/lobby"
	"elevator_control/elevator"
	"flag"
	"fmt"
	"time"
)

func main() {
	var clock clock.Clock
	clock.Cycle = 0

	var floors = flag.Int("floors", 5, "Number of floors the elevators will serve")
	flag.Parse()

	lobbyControlPanels := lobby.CreateControlPanels(*floors)
	fmt.Println(lobbyControlPanels)

	e1 := elevator.Elevator{
		ID:           1,
		CurrentFloor: 1,
		DestFloor:    1,
		DoorOpen:     false,
	}

	for {
		// start clock
		clockTimer := time.Now()

		//refresh
		fmt.Printf("[%v] ID: %d, Current: %d, Dest: %d, DoorOpen: %v\n", clock.Cycle, e1.ID, e1.CurrentFloor, e1.DestFloor, e1.DoorOpen)

		// force sync to clock
		clock.Step(&clockTimer)
	}

}
