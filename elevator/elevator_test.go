package elevator

import "testing"

// TestDoorOpen checks that opening a closed door will work
func TestDoorOpen(t *testing.T) {
	// given
	elevator1 := &Elevator{
		ID:           1,
		CurrentFloor: 1,
		DestFloor:    1,
		DoorOpen:     false,
	}

	// when
	_ = elevator1.OpenDoor()

	// then
	want := true
	got := (*elevator1).DoorOpen

	if want != got {
		t.Errorf("Door did not open. Want: %v, Have: %v", want, got)
	}

}

// TestDoorClose checks that closing a door will work
func TestDoorClose(t *testing.T) {
	// given
	elevator1 := &Elevator{
		ID:           1,
		CurrentFloor: 1,
		DestFloor:    1,
		DoorOpen:     true,
	}

	// when
	_ = elevator1.CloseDoor()

	// then
	want := false
	got := (*elevator1).DoorOpen

	if want != got {
		t.Errorf("Door did not close. Want: %v, Have: %v", want, got)
	}

}

func TestChangeFloor(t *testing.T) {
	// given
	elevator1 := &Elevator{
		ID:           1,
		CurrentFloor: 1,
		DestFloor:    1,
		DoorOpen:     true,
	}

	// when
	_ = elevator1.ChangeFloor(5)

	// then
	got := (*elevator1).CurrentFloor
	want := 5

	if want != got {
		t.Errorf("Door did not close. Want: %v, Have: %v", want, got)
	}
}
