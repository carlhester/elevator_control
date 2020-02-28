package lobby

func CreateControlPanels(floors int) []LobbyControl {
	LobbyControls := make([]LobbyControl, floors)
	for i := range LobbyControls {
		//if i != 0 {
		LobbyControls[i].ID = i + 1
		LobbyControls[i].Buttons = createControlPanelButtons(floors, i+1)
		//}
	}
	return LobbyControls
}

func createControlPanelButtons(floors int, curFloor int) []int {
	btns := []int{}
	for i := 1; i <= floors; i++ {
		if i != curFloor {
			btns = append(btns, i)
		}
	}
	return btns
}
