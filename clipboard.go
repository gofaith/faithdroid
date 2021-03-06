package faithdroid

type FClipboard struct {
	FBase
}

func Clipboard(a IActivity) *FClipboard {
	v := &FClipboard{}
	v.VID = NewToken()
	v.ClassName = "Clipboard"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FClipboard) OnChange(f func()) *FClipboard {
	fnId := NewToken()
	fn:=func(string) string {
		f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
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
