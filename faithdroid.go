package faithdroid

import (
	"github.com/gofaith/faithdroid/base"
	"github.com/gofaith/faithdroid/interfaces"
)

type UIInterface interface {
	interfaces.UIInterface
}

func Invoke(fnID int, args string) {
	base.Invoke(fnID, args)
}
