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
	GlobalVars.BaseMap[v.VID] = v
	return v
}
func (v *FSnackbar) text(s string) *FSnackbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FSnackbar) action(s string, onClick func()) *FSnackbar {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		onClick()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Action", JsonArray([]string{s, fnId}))
	return v
}
func (v *FSnackbar) show() {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Show", "")
}
func showSnackbar(a IActivity, s string) {
	Snackbar(a).text(s).show()
}
func showSnackbarWithAction(a IActivity, s string, act string, onClick func()) {
	Snackbar(a).text(s).action(act, onClick).show()
}
