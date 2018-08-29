package faithdroid

type FLinearLayout struct {
	FBaseView
	showAfter bool
	Children  []IView
}

func (v *FLinearLayout) setId(s string) *FLinearLayout {
	GlobalVars.idMap[s] = v
	return v
}
func getLinearLayoutById(id string) *FLinearLayout {
	if v, ok := GlobalVars.idMap[id].(*FLinearLayout); ok {
		return v
	}
	return nil
}
func linearlayout(a *Activity) *FLinearLayout {
	v := &FLinearLayout{}
	v.VID = newToken()
	v.ClassName = "LinearLayout"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	return v
}
func (v *FLinearLayout) size(w, h int) *FLinearLayout {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}

func (v *FLinearLayout) background(s string) *FLinearLayout {
	v.FBaseView.background(s)
	return v
}
func (v *FLinearLayout) backgroundColor(s int) *FLinearLayout {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FLinearLayout) cachedBackground(s string) *FLinearLayout {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FLinearLayout) onClick(f func()) *FLinearLayout {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
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
	v.FBaseView.layoutGravity(gravity)
	return v
}

func (v *FLinearLayout) elevation(dp float32) *FLinearLayout {
	v.FBaseView.elevation(dp)
	return v
}
func (v *FLinearLayout) assign(fb **FLinearLayout) *FLinearLayout {
	*fb = v
	return v
}

func (v *FLinearLayout) layoutWeight(f int) *FLinearLayout {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FLinearLayout) append(vs ...IView) *FLinearLayout {
	v.Children = vs
	for _, i := range vs {
		GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "AddView", i.getVID())
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
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Orientation", "VERTICAL")
	return v
}
func (v *FLinearLayout) horizontal() *FLinearLayout {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Orientation", "HORIZONTAL")
	return v
}
func (v *FLinearLayout) isVertical() bool {
	return GlobalVars.uis[v.UI].ViewGetAttr(v.VID, "Orientation") == "VERTICAL"
}
