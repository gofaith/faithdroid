package widget

import (
	"github.com/gofaith/faithdroid/interfaces"
	"github.com/gofaith/faithdroid/util"
	"github.com/gofaith/faithdroid/vars"
)

type FButton struct {
	FBaseView
}

func Button(w interfaces.IWindow) *FButton {
	v := &FButton{}
	v.Vid = util.NewNumToken()
	v.ClassName = vars.Button
	v.UI = w.GetUI()
	v.UI.New(v.ClassName, v.Vid)
	vars.SetView(v.Vid, v)
	return v
}
