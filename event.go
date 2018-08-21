package faithdroid

var (
	eventHandlersMap = make(map[string]func(string))
)

func TriggerEventHandler(eID string, args string) {
	if h, ok := eventHandlersMap[eID]; ok {
		h(args)
	}
}
