package vars

import "github.com/gofaith/faithdroid/interfaces"

type viewBinding struct {
	vid  int
	view interfaces.IView
}

var (
	viewMap     = make(map[int]interfaces.IView)
	viewMapChan = make(chan viewBinding, 1)
)

func init() {
	go func() {
		for bd := range viewMapChan {
			viewMap[bd.vid] = bd.view
		}
	}()
}

func GetView(vid int) interfaces.IView {
	return viewMap[vid]
}

func SetView(vid int, view interfaces.IView) {
	viewMapChan <- viewBinding{vid: vid, view: view}
}
