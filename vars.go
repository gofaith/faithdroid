package faithdroid

import "sync"

var GlobalVars = struct {
	IdMap                map[string]IView //only store views that user set special Id on
	ViewMap              map[string]IView //store every view
	EventHandlersMap     map[string]func(string) string
	EventHandlersMapLock sync.Mutex
	UIs                  map[string]UIController
}{
	make(map[string]IView),
	make(map[string]IView),
	make(map[string]func(string) string),
	sync.Mutex{},
	make(map[string]UIController),
}

func GetVidById(id string) string {
	if v, ok := GlobalVars.IdMap[id]; ok {
		return v.GetViewId()
	} else {
		panic(`id "` + id + `" doesn't exists`)
	}
}
