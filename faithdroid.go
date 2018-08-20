package faithdroid

type Activity struct {
	ui UIInterface
}

func (a *Activity) SetUIInterface(u UIInterface) {
	a.ui = u
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
