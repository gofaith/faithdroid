package faithdroid

type FService struct {
}

func (s *FService) GetContext() string {
	return "Service"
}
func StartService(a IActivity, f func()) *FService {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		go f()
		return ""
	}
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Service", "OnCreate", fnId)
	return &FService{}
}
