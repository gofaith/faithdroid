package faithdroid

type FAlertDialog struct {
	FBaseView
}

func AlertDialog(a *Activity) *FAlertDialog {
	v := &FAlertDialog{}
	v.VID = NewToken()
	v.ClassName = "AlertDialog"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}

func (v *FAlertDialog) Title(s string) *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FAlertDialog) View(iv IView) *FAlertDialog {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "View", iv.GetViewId())
	return v
}
func (v *FAlertDialog) PositiveButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		onClick(v)
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PositiveButton", JsonArray([]string{text, fnId}))
	return v
}

func (v *FAlertDialog) NegativeButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		onClick(v)
		return ""
	}
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
