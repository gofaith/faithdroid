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
type FBase struct {
	VID, ClassName string
	UI             string
}
type FBaseView struct {
	FBase
	Fn_TopToTop func()
	Fn_TopToBottom func()
	Fn_BottomToBottom func()
	Fn_BottomToTop func()
	Fn_LeftToLeft func()
	Fn_LeftToRight func()
	Fn_RightToRight func()
	Fn_RightToLeft func()
}
type IBase interface {
	GetViewId() string
}
type IBaseView interface {
	IBase
	GetBaseView()*FBaseView
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
func (v *FBaseView) Foreground(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go DownloadNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Foreground", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Foreground", s)
}
func (v *FBaseView) CachedForeground(s string) {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go CacheNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Foreground", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunOnUIThread(fnID)
		})
		return
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Foreground", s)
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

func (v *FBaseView) TopToTop(iv IBaseView) {
	v.Fn_TopToTop= func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TopToTop", vid)
	}
}
func (v *FBaseView) TopToBottom(iv IBaseView) {
	v.Fn_TopToBottom=func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TopToBottom", vid)
	}
}
func (v *FBaseView) BottomToBottom(iv IBaseView) {
	v.Fn_BottomToBottom = func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "BottomToBottom", vid)
	}
}
func (v *FBaseView) BottomToTop(iv IBaseView) {
	v.Fn_BottomToTop = func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "BottomToTop", vid)
	}
}
func (v *FBaseView) LeftToLeft(iv IBaseView) {
	v.Fn_LeftToLeft=func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "LeftToLeft", vid)
	}
}
func (v *FBaseView) LeftToRight(iv IBaseView) {
	v.Fn_LeftToRight=func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "LeftToRight", vid)
	}
}
func (v *FBaseView) RightToRight(iv IBaseView) {
	v.Fn_RightToRight=func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "RightToRight", vid)
	}
}
func (v *FBaseView) RightToLeft(iv IBaseView) {
	v.Fn_RightToLeft= func() {
		vid:= "Parent"
		if iv != nil {
			vid=iv.GetViewId()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "RightToLeft", vid)
	}
}
