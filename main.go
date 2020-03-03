package main

import (
	"fmt"
)

type elevator struct {
	ID           int
	CurrentFloor int
	DestFloor    int
}

type floor struct {
	level int
}

func NewFloor(level int) *floor {
	floor := floor{level: level}
	return &floor
}

func NewElevator(id int, CurrentFloor int) *elevator {
	elevator := elevator{
		ID:           id,
		CurrentFloor: CurrentFloor,
		DestFloor:    CurrentFloor,
	}
	return &elevator
}

func (e *elevator) showStatus() {
	fmt.Printf("Elevator ID: %d.  Current: %d.  Destination: %d.\n", e.ID, e.CurrentFloor, e.DestFloor)
}

func (e *elevator) move(dst int) {
	e.DestFloor = dst
	for e.CurrentFloor != dst {
		if e.CurrentFloor > dst {
			e.CurrentFloor -= 1
			fmt.Printf("Elevator %d going down.\n", e.ID)
		} else if e.CurrentFloor < dst {
			e.CurrentFloor += 1
			fmt.Printf("Elevator %d going up.\n", e.ID)
		}
	}
	return
}

func main() {
	var floors [5]*floor
	floors[0] = NewFloor(1)
	floors[1] = NewFloor(2)
	floors[2] = NewFloor(3)
	floors[3] = NewFloor(4)
	floors[4] = NewFloor(5)

	e1 := *NewElevator(1, floors[0].level)
	e1.showStatus()
	e1.move(3)
	e1.showStatus()
	e1.move(1)
	e1.showStatus()
	e1.move(5)
	e1.showStatus()

}
