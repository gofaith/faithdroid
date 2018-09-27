package faithdroid

var GlobalVars = struct {
	IdMap            map[string]IBaseView //only store views that user set special Id on
	BaseMap          map[string]IBase //store every view
	EventHandlersMap map[string]func(string) string
	UIs              map[string]UIController
}{
	make(map[string]IBaseView),
	make(map[string]IBase),
	make(map[string]func(string) string),
	make(map[string]UIController),
}
