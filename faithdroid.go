package faithdroid

import (
	"github.com/gofaith/faithdroid/base"
	"github.com/gofaith/faithdroid/interfaces"
)

type (
	UIBridge interface {
		interfaces.UIBridge
	}
	Window struct {
		ui                      interfaces.UIBridge
		fnOnCreate, fnOnDestroy func()
	}
	MainWindow struct {
		Window
	}
)

func (w *Window) OnCreate() {
	if w.fnOnCreate != nil {
		w.fnOnCreate()
	}
}
func (w *Window) OnDestroy() {
	if w.fnOnDestroy != nil {
		w.fnOnDestroy()
	}
}

func (w *Window) GetUI() interfaces.UIBridge {
	return w.ui
}

func (w *Window) SetUI(ui UIBridge) {
	w.ui = ui
}

func Invoke(fnID int, args string) {
	base.Invoke(fnID, args)
}
