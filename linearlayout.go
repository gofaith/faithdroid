package faithdroid

type FLinearLayout struct {
	FBaseView
	showAfter bool
}

func (v *FLinearLayout) getVID() string {
	return v.vID
}
func (v *FLinearLayout) setId(s string) *FLinearLayout {
	idMap[s] = v
	return v
}
func getLinearLayoutById(id string) *FLinearLayout {
	if v, ok := idMap[id].(*FLinearLayout); ok {
		return v
	}
	return nil
}
func linearlayout(a *Activity) *FLinearLayout {
	v := &FLinearLayout{}
	v.vID = newToken()
	v.className = "LinearLayout"
	v.ui = a.ui
	a.ui.NewView(v.className, v.vID)
	return v
}
func (v *FLinearLayout) size(w, h int) *FLinearLayout {
	i := []int{w, h}
	v.ui.ViewSetAttr(v.vID, "Size", jsonArray(i))
	return v
}

func (v *FLinearLayout) background(s string) *FLinearLayout {
	v.FBaseView.background(s)
	return v
}
func (v *FLinearLayout) backgroundColor(s string) *FLinearLayout {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FLinearLayout) cachedBackground(s string) *FLinearLayout {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FLinearLayout) onClick(f func()) *FLinearLayout {
	fnID := newToken()
	eventHandlersMap[fnID] = func(string) {
		f()
	}
	v.ui.ViewSetAttr(v.vID, "OnClick", fnID)
	return v
}

func (v *FLinearLayout) visible() *FLinearLayout {
	v.FBaseView.visible()
	return v
}
func (v *FLinearLayout) invisible() *FLinearLayout {
	v.FBaseView.invisible()
	return v
}
func (v *FLinearLayout) gone() *FLinearLayout {
	v.FBaseView.gone()
	return v
}

func (v *FLinearLayout) padding(left, top, right, bottom int) *FLinearLayout {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FLinearLayout) paddingLeft(dp int) *FLinearLayout {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FLinearLayout) paddingTop(dp int) *FLinearLayout {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FLinearLayout) paddingRight(dp int) *FLinearLayout {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FLinearLayout) paddingBottom(dp int) *FLinearLayout {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FLinearLayout) paddingAll(all int) *FLinearLayout {
	v.FBaseView.padding(all, all, all, all)
	return v
}

func (v *FLinearLayout) margin(left, top, right, bottom int) *FLinearLayout {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FLinearLayout) marginLeft(dp int) *FLinearLayout {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FLinearLayout) marginTop(dp int) *FLinearLayout {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FLinearLayout) marginRight(dp int) *FLinearLayout {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FLinearLayout) marginBottom(dp int) *FLinearLayout {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FLinearLayout) marginAll(dp int) *FLinearLayout {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FLinearLayout) layoutGravity(gravity int) *FLinearLayout {
	v.ui.ViewSetAttr(v.vID, "LayoutGravity", sPrintf("%v", gravity))
	return v
}

// --------------------------------------------------------
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
