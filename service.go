package faithdroid

type FService struct {
}

func (s *FService) GetContext() string {
	return "Service"
}
func StartService(a *Activity, f func()) *FService {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		go f()
		return ""
	}
	GlobalVars.UIs[a.UI].ViewSetAttr("Service", "OnCreate", fnId)
	return &FService{}
}
