package faithdroid

type FClipboard struct {
	FBaseView
}

func Clipboard(a *Activity) *FClipboard {
	v := &FClipboard{}
	v.VID = NewToken()
	v.ClassName = "Clipboard"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.BaseMap[v.VID] = v
	return v
}
func (v *FClipboard) OnChange(f func()) *FClipboard {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnChange", fnId)
	return v
}
func (v *FClipboard) ClipData(s string) *FClipboard {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ClipData", s)
	return v
}
func (v *FClipboard) GetClipCount() int {
	i, e := a2i(GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ClipCount"))
	if e != nil {
		return 0
	}
	return i
}
func (v *FClipboard) GetClipData() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ClipData")
}
func (v *FClipboard) GetClipDataAt(index int) string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Item"+SPrintf(index))
}
