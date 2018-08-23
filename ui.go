package faithdroid

type UIInterface interface {
	NewView(viewName string, vID string)
	ViewSetAttr(vID string, attr string, value string)
	ViewGetAttr(vID string, attr string) string
	ShowOnRootView(vID string)
	RunOnUIThread(fnID string)
	GetPkg() string
}
type FBaseView struct {
	vID, className string
	ui             UIInterface
}
type IView interface {
	getVID() string
}

func (v *FBaseView) show() {
	v.ui.ShowOnRootView(v.vID)
}
func (v *FBaseView) background(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go downloadNetFile(s, "/data/data/"+v.ui.GetPkg()+"/tmp/", func(f string) {
			fnID := newToken()
			eventHandlersMap[fnID] = func(string) {
				v.ui.ViewSetAttr(v.vID, "Background", "file://"+f)
			}
			v.ui.RunOnUIThread(fnID)
		})
		return
	}
	v.ui.ViewSetAttr(v.vID, "Background", s)
}
func (v *FBaseView) cachedBackground(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go cacheNetFile(s, "/data/data/"+v.ui.GetPkg()+"/cacheDir/", func(f string) {
			fnID := newToken()
			eventHandlersMap[fnID] = func(string) {
				v.ui.ViewSetAttr(v.vID, "Background", "file://"+f)
			}
			v.ui.RunOnUIThread(fnID)
		})
		return
	}
	v.ui.ViewSetAttr(v.vID, "Background", s)
}
func (v *FBaseView) backgroundColor(s string) {
	v.ui.ViewSetAttr(v.vID, "BackgroundColor", s)
}
func (v *FBaseView) visible() {
	v.ui.ViewSetAttr(v.vID, "Visibility", "VISIBLE")
}
func (v *FBaseView) invisible() {
	v.ui.ViewSetAttr(v.vID, "Visibility", "INVISIBLE")
}
func (v *FBaseView) gone() {
	v.ui.ViewSetAttr(v.vID, "Visibility", "GONE")
}
func (v *FBaseView) padding(left, top, right, bottom int) {
	v.ui.ViewSetAttr(v.vID, "Padding", jsonArray([]int{left, top, right, bottom}))
}

var (
	idMap = make(map[string]IView)
)
