package faithdroid

type Activity struct {
	UI string
}

func (a *Activity) SetUIInterface(u UIController) string {
	uId := newToken()
	GlobalVars.uis[uId] = u
	a.UI = uId
	return uId
}
func (a *Activity) OnCreate() {

}
func (a *Activity) OnResume() {
}
func (a *Activity) OnPause() {
}
func (a *Activity) OnStart() {
}
func (a *Activity) OnDestroy() {
}

type ActivityConfig struct {
	LaunchMode string
	FnId       string
}

func (a *Activity) startActivity(createView func(*Activity), conf ActivityConfig) {
	fnId := newToken()
	GlobalVars.eventHandlersMap[fnId] = func(uId string) string {
		createView(&Activity{UI: uId})
		return ""
	}
	if conf.LaunchMode == "" {
		conf.LaunchMode = "Standard"
	}
	conf.FnId = fnId
	GlobalVars.uis[a.UI].NewView("Activity", jsonObject(conf))
}
func (a *Activity) finish() {
	GlobalVars.uis[a.UI].ViewSetAttr("Activity", "Finish", "")
}
