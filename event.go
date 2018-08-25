package faithdroid

func TriggerEventHandler(eID string, args string) string {
	if h, ok := GlobalVars.eventHandlersMap[eID]; ok {
		return h(args)
	}
	return ""
}
