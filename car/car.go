package car

import "log"

import "elevator/display"
import "elevator/control"

func NewCar(floors []int) (Car, error) {
	var car Car
	var err error
	car.ID = 1
	car.CurrentFloor = 1

	car.ControlPanel, err = control.NewCarControls(floors)
	if err != nil {
		log.Println("Error creating Car Controls", err)
	}

	car.Display, err = display.NewCarDisplay()
	if err != nil {
		log.Println("Error creating Car Display", err)
	}
	return car, nil

}
