package faithdroid

import (
	"github.com/gofaith/faithdroid/widget"
)

func (w *MainWindow) OnCreate() {
	widget.Button(w).Show()
}
