package faithdroid

var GlobalVars = struct {
	IdMap            map[string]IView //only store views that user set special Id on
	ViewMap          map[string]IView //store every view
	EventHandlersMap map[string]func(string) string
	UIs              map[string]UIController
}{
	make(map[string]IView),
	make(map[string]IView),
	make(map[string]func(string) string),
	make(map[string]UIController),
}
