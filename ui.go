package faithdroid

type UIController interface {
	NewView(viewName string, VID string)
	ViewSetAttr(VID string, attr string, value string)
	ViewGetAttr(VID string, attr string) string
	ShowOnRootView(VID string)
	RunOnUIThread(fnID string)
	GetPkg() string
}
type FBaseView struct {
	VID, ClassName string
	UI             string
}
type IView interface {
	getVID() string
}

func (v FBaseView) getVID() string {
	return v.VID
}
func (v *FBaseView) show() {
	GlobalVars.uis[v.UI].ShowOnRootView(v.VID)
}
func (v *FBaseView) background(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go downloadNetFile(s, "/data/data/"+GlobalVars.uis[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := newToken()
			GlobalVars.eventHandlersMap[fnID] = func(string) string {
				GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Background", "file://"+f)
				return ""
			}
			GlobalVars.uis[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Background", s)
}
func (v *FBaseView) cachedBackground(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go cacheNetFile(s, "/data/data/"+GlobalVars.uis[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := newToken()
			GlobalVars.eventHandlersMap[fnID] = func(string) string {
				GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Background", "file://"+f)
				return ""
			}
			GlobalVars.uis[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Background", s)
}
func (v *FBaseView) backgroundColor(s int) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "BackgroundColor", "#"+xPrintf(s))
}
func (v *FBaseView) visible() {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Visibility", "VISIBLE")
}
func (v *FBaseView) invisible() {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Visibility", "INVISIBLE")
}
func (v *FBaseView) gone() {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Visibility", "GONE")
}
func (v *FBaseView) padding(left, top, right, bottom int) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Padding", jsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) margin(left, top, right, bottom int) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Margin", jsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) layoutGravity(gravity int) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "LayoutGravity", sPrintf(gravity))
}
func (v *FBaseView) elevation(dp float32) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Elevation", sPrintf(dp))
}
func (v *FBaseView) layoutWeight(f int) {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "LayoutWeight", sPrintf(f))
}
