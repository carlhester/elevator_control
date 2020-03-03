package control

func NewCarControls(floors []int) (CarControlPanel, error) {
	var control CarControlPanel
	control.ID = 1
	control.FloorButtons = floors

	return control, nil
}
