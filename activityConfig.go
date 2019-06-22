package faithdroid

type ActivityConfig struct {
	MyLaunchMode string `json:"myLaunchMode"`
	MyFnId       string `json:"myFnId"`
	MyIntent     Intent `json:"myIntent"`
}

func NewActivityConfig() *ActivityConfig {
	ac := &ActivityConfig{}
	ac.MyIntent.Extras = make(map[string]string)
	return ac
}
func (c *ActivityConfig) LaunchMode_Standard() *ActivityConfig {
	c.MyLaunchMode = "Standard"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleTask() *ActivityConfig {
	c.MyLaunchMode = "SingleTask"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleInstance() *ActivityConfig {
	c.MyLaunchMode = "SingleInstance"
	return c
}
func (c *ActivityConfig) LaunchMode_SingleTop() *ActivityConfig {
	c.MyLaunchMode = "SingleTop"
	return c
}
func (c *ActivityConfig) IntentAction(s string) *ActivityConfig {
	c.MyIntent.Action = s
	return c
}
func (c *ActivityConfig) Paths(ps ...string) *ActivityConfig {
	c.MyIntent.Paths = ps
	return c
}
func (c *ActivityConfig) PutExtra(key, value string) *ActivityConfig {
	c.MyIntent.Extras[key] = value
	return c
}
func StartActivity(a IActivity, createView func(IActivity), conf *ActivityConfig) {
	fnId := NewToken()
	fn:=func(uId string) string {
		if createView != nil {
			createView(GlobalVars.UIs[uId].GetCurrentFActivity())
		}
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	if conf == nil {
		conf = NewActivityConfig()
	}
	conf.MyFnId = fnId
	GlobalVars.UIs[a.GetMyActivity().UI].NewView("Activity", JsonObject(conf))
}
func StartUriIntent(a IActivity, uri string) {
	GlobalVars.UIs[a.GetMyActivity().UI].ViewSetAttr("Activity", "StartUriIntent", uri)
}
