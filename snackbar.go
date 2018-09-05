package faithdroid

type FSnackbar struct {
	FBaseView
}

func Snackbar(a *Activity) *FSnackbar {
	v := &FSnackbar{}
	v.VID = newToken()
	v.ClassName = "Snackbar"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (v *FSnackbar) text(s string) *FSnackbar {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FSnackbar) action(s string, onClick func()) *FSnackbar {
	fnId := newToken()
	GlobalVars.eventHandlersMap[fnId] = func(string) string {
		onClick()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Action", jsonArray([]string{s, fnId}))
	return v
}
func (v *FSnackbar) show() {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Show", "")
}
func showSnackbar(a *Activity, s string) {
	Snackbar(a).text(s).show()
}
func showSnackbarWithAction(a *Activity, s string, act string, onClick func()) {
	Snackbar(a).text(s).action(act, onClick).show()
}
