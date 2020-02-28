package elevator

type Elevator struct {
	ID           int
	Buttons      []int
	CurrentFloor int
	DestFloor    int
	DoorOpen     bool
}
