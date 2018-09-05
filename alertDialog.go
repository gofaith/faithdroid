package faithdroid

type FAlertDialog struct {
	FBaseView
}

func AlertDialog(a *Activity) *FAlertDialog {
	v := &FAlertDialog{}
	v.VID = newToken()
	v.ClassName = "AlertDialog"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}

func (v *FAlertDialog) title(s string) *FAlertDialog {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FAlertDialog) view(iv IView) *FAlertDialog {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "View", iv.getVID())
	return v
}
func (v *FAlertDialog) positiveButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := newToken()
	GlobalVars.eventHandlersMap[fnId] = func(string) string {
		onClick(v)
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "PositiveButton", jsonArray([]string{text, fnId}))
	return v
}

func (v *FAlertDialog) negativeButton(text string, onClick func(*FAlertDialog)) *FAlertDialog {
	fnId := newToken()
	GlobalVars.eventHandlersMap[fnId] = func(string) string {
		onClick(v)
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "NegativeButton", jsonArray([]string{text, fnId}))
	return v
}
func (v *FAlertDialog) show() *FAlertDialog {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Show", "")
	return v
}
func (v *FAlertDialog) dismiss() *FAlertDialog {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Dismiss", "")
	return v
}
