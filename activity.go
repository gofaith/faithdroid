package faithdroid

type Activity struct {
	UI string
}

func (a *Activity) SetUIInterface(u UIController) {
	uId := newToken()
	GlobalVars.uis[uId] = u
	a.UI = uId
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
