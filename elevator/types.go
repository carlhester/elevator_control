package elevator

import "elevator_control/controls/car"

type Elevator struct {
	ID           int
	CurrentFloor int
	DestFloor    int
	DoorOpen     bool
	Buttons      car.CarControl
	Display      car.CarDisplay
}
