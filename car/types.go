package car

import "elevator/display"
import "elevator/control"

type Car struct {
	ID               int
	ControlPanel     control.CarControlPanel
	Display          display.CarDisplay
	InService        bool
	CurrentFloor     int
	DestinationFloor int
}
