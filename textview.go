package faithdroid

type FTextView struct {
	FBaseView
}

func (v *FTextView) getVID() string {
	return v.vID
}
func textview(a *Activity) *FTextView {
	v := &FTextView{}
	v.vID = newToken()
	v.className = "TextView"
	v.ui = a.ui
	a.ui.NewView(v.className, v.vID)
	return v
}
func (v *FTextView) text(s string) *FTextView {
	v.ui.ViewSetAttr(v.vID, "Text", s)
	return v
}
func (v *FTextView) getText() string {
	return v.ui.ViewGetAttr(v.vID, "Text")
}
