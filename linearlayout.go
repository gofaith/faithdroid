package faithdroid

type FLinearLayout struct {
	FBaseView
}

func linearlayout(a *Activity) *FLinearLayout {
	vID := newToken()
	v := &FLinearLayout{}
	v.vID = vID
	v.className = "LinearLayout"
	v.ui = a.ui
	a.ui.NewView(v.className, v.vID)
	return v
}

func (v *FLinearLayout) Append(vs ...IView) *FLinearLayout {
	for _, i := range vs {
		v.ui.ViewSetAttr(v.vID, "AddView", i.getVID())
	}
	return v
}
