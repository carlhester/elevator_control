package elevator

func (e *Elevator) OpenDoor() error {
	if e.DoorOpen == true {
		return nil //door already open error
	}

	e.DoorOpen = true
	return nil
}

func (e *Elevator) CloseDoor() error {
	if e.DoorOpen == false {
		return nil //door already closed error
	}

	e.DoorOpen = false
	return nil
}

func (e *Elevator) ChangeFloor(requestFloor int) error {
	e.CurrentFloor = requestFloor
	return nil
}
