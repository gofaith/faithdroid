package faithdroid

type FLinearLayout struct {
	FBaseView
	showAfter bool
}

func (v *FLinearLayout) getVID() string {
	return v.vID
}
func linearlayout(a *Activity) *FLinearLayout {
	v := &FLinearLayout{}
	v.vID = newToken()
	v.className = "LinearLayout"
	v.ui = a.ui
	a.ui.NewView(v.className, v.vID)
	return v
}

func (v *FLinearLayout) append(vs ...IView) *FLinearLayout {
	for _, i := range vs {
		v.ui.ViewSetAttr(v.vID, "AddView", i.getVID())
	}
	if v.showAfter {
		v.show()
	}
	return v
}
func (v *FLinearLayout) deferShow() *FLinearLayout {
	v.showAfter = true
	return v
}
func (v *FLinearLayout) vertical() *FLinearLayout {
	v.ui.ViewSetAttr(v.vID, "Orientation", "VERTICAL")
	return v
}
func (v *FLinearLayout) horizontal() *FLinearLayout {
	v.ui.ViewSetAttr(v.vID, "Orientation", "HORIZONTAL")
	return v
}
func (v *FLinearLayout) isVertical() bool {
	return v.ui.ViewGetAttr(v.vID, "Orientation") == "VERTICAL"
}
