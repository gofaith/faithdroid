package faithdroid

type FButton struct {
	FBaseView
}

func (v *FButton) getVID() string {
	return v.vID
}
func button(a *Activity) *FButton {
	v := &FButton{}
	v.vID = newToken()
	v.className = "Button"
	v.ui = a.ui
	v.ui.NewView(v.className, v.vID)
	return v
}
func (v *FButton) size(w, h int) *FButton {
	i := []int{w, h}
	v.ui.ViewSetAttr(v.vID, "Size", jsonArray(i))
	return v
}
func (v *FButton) text(s string) *FButton {
	v.ui.ViewSetAttr(v.vID, "Text", s)
	return v
}
func (v *FButton) getText() string {
	return v.ui.ViewGetAttr(v.vID, "Text")
}
func (v *FButton) enabled(b bool) *FButton {
	if b {
		v.ui.ViewSetAttr(v.vID, "Enabled", "true")
	} else {
		v.ui.ViewSetAttr(v.vID, "Enabled", "false")
	}
	return v
}
func (v *FButton) isEnabled() bool {
	if v.ui.ViewGetAttr(v.vID, "Enabled") == "true" {
		return true
	}
	return false
}
