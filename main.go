package main

import (
	"elevator_control/controls/lobby"
	"elevator_control/elevator"
	"flag"
	"fmt"
)

func main() {

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

	fmt.Println(e1)
}
