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

var (
	idMap = make(map[string]IView)
)
