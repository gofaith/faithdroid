package faithdroid

var GlobalVars = struct {
	idMap            map[string]IView //only store views that user set special Id on
	viewMap          map[string]IView //store every view
	eventHandlersMap map[string]func(string) string
	uis              map[string]UIController
}{
	make(map[string]IView),
	make(map[string]IView),
	make(map[string]func(string) string),
	make(map[string]UIController),
}
