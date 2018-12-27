package faithdroid

type FService struct {
	a IActivity
}

func (s *FService) GetContext() string {
	return "Service"
}
func StartService(a IActivity, f func()) *FService {
	fnId := NewToken()
	fn := func(string) string {
		go f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Service", "OnCreate", fnId)
	return &FService{}
}
func Service(a IActivity) *FService {
	return &FService{a: a}
}
func (f *FService) QuitButton(title string) *FService {
	GlobalVars.UIs[f.a.GetMyActivity().UI].ViewSetAttr("Service", "QuitButton", title)
	return f
}
func (v *FService) Title(s string) *FService {
	GlobalVars.UIs[v.a.GetMyActivity().UI].ViewSetAttr("Service", "Title", s)
	return v
}
func (v *FService) SubTitle(s string) *FService {
	GlobalVars.UIs[v.a.GetMyActivity().UI].ViewSetAttr("Service", "SubTitle", s)
	return v
}
func (v *FService) OnCreate(f func()) *FService {
	fnID := NewToken()
	fn := func(string) string {
		f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnID] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.a.GetMyActivity().UI].ViewSetAttr("Service", "OnCreate", fnID)
	return v
}
func (v *FService) OnQuit(f func()) *FService {
	fnID := NewToken()
	fn := func(string) string {
		f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnID] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.a.GetMyActivity().UI].ViewSetAttr("Service", "OnQuit", fnID)
	return v
}
func (f *FService) Show() *FService {
	GlobalVars.UIs[f.a.GetMyActivity().UI].ViewSetAttr("Service", "Show", "")
	return f
}
func (f *FService) FinishAllActivity() *FService {
	GlobalVars.UIs[f.a.GetMyActivity().UI].ViewSetAttr("Service", "FinishAllActivity", "")
	return f
}
func (f *FService) KillAll() *FService {
	GlobalVars.UIs[f.a.GetMyActivity().UI].ViewSetAttr("Service", "KillAll", "")
	return f
}
func (f *FService) Stop() *FService {
	GlobalVars.UIs[f.a.GetMyActivity().UI].ViewSetAttr("Service", "Stop", "")
	return f
}
