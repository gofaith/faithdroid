package faithdroid

func TriggerEventHandler(eID string, args string) string {
	if h, ok := GlobalVars.EventHandlersMap[eID]; ok {
		return h(args)
	}
	return ""
}
