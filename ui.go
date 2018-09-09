package faithdroid

type UIController interface {
	NewView(viewName string, VID string)
	ViewSetAttr(VID string, attr string, value string)
	ViewGetAttr(VID string, attr string) string
	ShowOnRootView(VID string)
	RunOnUIThread(fnID string)
	GetPkg() string
	GetCurrentFActivity() *Activity
}
type FBaseView struct {
	FBase
}
type FBase struct {
	VID, ClassName string
	UI             string
}
type IView interface {
	GetViewId() string
}

func (v FBase) GetViewId() string {
	return v.VID
}
func (v *FBaseView) Show() {
	GlobalVars.UIs[v.UI].ShowOnRootView(v.VID)
}
func (v *FBaseView) Background(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go DownloadNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Background", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Background", s)
}
func (v *FBaseView) CachedBackground(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go CacheNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Background", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Background", s)
}
func (v *FBaseView) BackgroundColor(s string) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "BackgroundColor", s)
}
func (v *FBaseView) Visible() {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Visibility", "VISIBLE")
}
func (v *FBaseView) Invisible() {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Visibility", "INVISIBLE")
}
func (v *FBaseView) Gone() {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Visibility", "GONE")
}
func (v *FBaseView) Padding(left, top, right, bottom int) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Padding", JsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) Margin(left, top, right, bottom int) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Margin", JsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) LayoutGravity(gravity int) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "LayoutGravity", SPrintf(gravity))
}
func (v *FBaseView) Elevation(dp float64) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Elevation", SPrintf(dp))
}
func (v *FBaseView) LayoutWeight(f int) {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "LayoutWeight", SPrintf(f))
}
