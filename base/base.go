package base

import (
	"github.com/gofaith/faithdroid/interfaces"
)

type FBase struct {
	Vid       int
	ClassName int
	UI        interfaces.UIBridge
}

func (f FBase) GetVid() int {
	return f.Vid
}
