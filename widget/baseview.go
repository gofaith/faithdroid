package widget

import (
	"github.com/gofaith/faithdroid/base"
)

type FBaseView struct {
	base.FBase
}

func (f *FBaseView) Show() {
	f.UI.Show(f.Vid)
}
