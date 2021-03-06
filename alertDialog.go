package faithdroid

type FAlertDialog struct {
	FBaseView
}

func AlertDialog(a IActivity) *FAlertDialog {
	v := &FAlertDialog{}
	v.VID = NewToken()
	v.ClassName = "AlertDialog"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}

func (v *FAlertDialog) Title(s string) *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FAlertDialog) Append(iv IView) *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "View", iv.GetViewId())
	return v
}
func (v *FAlertDialog) PositiveButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := NewToken()
	fn:=func(string) string {
		onClick(v)
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PositiveButton", JsonArray([]string{text, fnId}))
	return v
}

func (v *FAlertDialog) NegativeButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := NewToken()
	fn:=func(string) string {
		onClick(v)
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "NegativeButton", JsonArray([]string{text, fnId}))
	return v
}
func (v *FAlertDialog) Show() *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Show", "")
	return v
}
func (v *FAlertDialog) Dismiss() *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Dismiss", "")
	return v
}
