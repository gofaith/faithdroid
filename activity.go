package faithdroid

type Activity struct {
	UI string
}

func (a *Activity) SetUIInterface(u UIController) string {
	uId := NewToken()
	GlobalVars.UIs[uId] = u
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

func StartActivity(a *Activity, createView func(*Activity), conf *ActivityConfig) {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(uId string) string {
		if createView != nil {
			createView(&Activity{UI: uId})
		}
		return ""
	}
	if conf == nil {
		conf = &ActivityConfig{}
		conf.LaunchMode = "Standard"
	}
	conf.FnId = fnId
	GlobalVars.UIs[a.UI].NewView("Activity", JsonObject(conf))
}
func Finish(a *Activity) {
	GlobalVars.UIs[a.UI].ViewSetAttr("Activity", "Finish", "")
}
func ShowToast(a *Activity, s string) {
	GlobalVars.UIs[a.UI].ViewSetAttr("Activity", "ShowToast", s)
}
