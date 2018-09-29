package faithdroid

type FSnackbar struct {
	FBaseView
}

func Snackbar(a IActivity) *FSnackbar {
	v := &FSnackbar{}
	v.VID = NewToken()
	v.ClassName = "Snackbar"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FSnackbar) Text(s string) *FSnackbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FSnackbar) Action(s string, onClick func()) *FSnackbar {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		onClick()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Action", JsonArray([]string{s, fnId}))
	return v
}
func (v *FSnackbar) Show() {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Show", "")
}
func ShowSnackbar(a IActivity, s string) {
	Snackbar(a).Text(s).Show()
}
func ShowSnackbarWithAction(a IActivity, s string, act string, onClick func()) {
	Snackbar(a).Text(s).Action(act, onClick).Show()
}
