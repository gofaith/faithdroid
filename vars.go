package faithdroid

var GlobalVars = struct {
	idMap            map[string]IView
	eventHandlersMap map[string]func(string) string
	uis              map[string]UIController
}{
	make(map[string]IView),
	make(map[string]func(string) string),
	make(map[string]UIController),
}
