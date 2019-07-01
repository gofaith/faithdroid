package base

import (
	"github.com/gofaith/faithdroid/interfaces"
	"github.com/gofaith/faithdroid/util"
	"github.com/gofaith/faithdroid/vars"
)

func Invoke(fnID int, args string) {
	vars.GetFn(fnID)(args)
}

func RunOnUIThread(ui interfaces.UIBridge, fn func()) {
	fnID := util.NewNumToken()
	vars.SetFn(fnID, func(string) string {
		fn()
		return ""
	})
	ui.RunOnUIThread(fnID)
}
