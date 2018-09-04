package faithdroid

type FButton struct {
	FBaseView
}

func Button(a *Activity) *FButton {
	v := &FButton{}
	v.VID = newToken()
	v.ClassName = "Button"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) getButtonByItemId(id string) *FButton {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FButton); ok {
			return bt
		}
	}
	return nil
}

func (v *FButton) size(w, h int) *FButton {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FButton) setId(s string) *FButton {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FButton) setItemId(parent *FListView, id string) *FButton {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getButtonById(id string) *FButton {
	if v, ok := GlobalVars.idMap[id].(*FButton); ok {
		return v
	}
	return nil
}

func (v *FButton) background(s string) *FButton {
	v.FBaseView.background(s)
	return v
}
func (v *FButton) backgroundColor(s int) *FButton {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FButton) cachedBackground(s string) *FButton {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FButton) visible() *FButton {
	v.FBaseView.visible()
	return v
}
func (v *FButton) invisible() *FButton {
	v.FBaseView.invisible()
	return v
}
func (v *FButton) gone() *FButton {
	v.FBaseView.gone()
	return v
}

func (v *FButton) padding(left, top, right, bottom int) *FButton {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FButton) paddingLeft(dp int) *FButton {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FButton) paddingTop(dp int) *FButton {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FButton) paddingRight(dp int) *FButton {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FButton) paddingBottom(dp int) *FButton {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FButton) paddingAll(dp int) *FButton {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FButton) margin(left, top, right, bottom int) *FButton {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FButton) marginLeft(dp int) *FButton {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FButton) marginTop(dp int) *FButton {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FButton) marginRight(dp int) *FButton {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FButton) marginBottom(dp int) *FButton {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FButton) marginAll(dp int) *FButton {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FButton) layoutGravity(gravity int) *FButton {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FButton) elevation(dp float32) *FButton {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FButton) assign(fb **FButton) *FButton {
	*fb = v
	return v
}
func (v *FButton) layoutWeight(f int) *FButton {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FButton) text(s string) *FButton {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FButton) textColor(s string) *FButton {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}
func (v *FButton) textSize(dpsize int) *FButton {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextSize", sPrintf(dpsize))
	return v
}
func (v *FButton) getText() string {
	return GlobalVars.uis[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FButton) enabled(b bool) *FButton {
	if b {
		GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FButton) isEnabled() bool {
	if GlobalVars.uis[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}

func (v *FButton) onClick(f func()) *FButton {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
