package lobby

import "testing"

func TestCreateControlPanels(t *testing.T) {
	// given
	floors := 5

	// when
	lobbyControlPanels := CreateControlPanels(floors)

	// then
	want := 5
	got := len(lobbyControlPanels)

	if want != got {
		t.Errorf("Did not create correct number of panels")
	}
}

func TestCreateControlPanelButtons(t *testing.T) {
	// given
	floors := 5

	// when
	btns := createControlPanelButtons(floors, 3)

	// then
	want := 4
	got := len(btns)

	if want != got {
		t.Errorf("Did not create correct number of btns.  Want: %v, Got: %v", want, got)
	}
}
