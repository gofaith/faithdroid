package vars

import (
	"github.com/gofaith/faithdroid/interfaces"
)

var (
	uiMap = make(map[int]interfaces.UIBridge)
)

func GetUI(id int) interfaces.UIBridge {
	return uiMap[id]
}

func SetUI(id int, ui interfaces.UIBridge) {
	uiMap[id] = ui
}
