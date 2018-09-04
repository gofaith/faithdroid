package faithdroid

type FTextView struct {
	FBaseView
}

func TextView(a *Activity) *FTextView {
	v := &FTextView{}
	v.VID = newToken()
	v.ClassName = "TextView"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) getTextViewByItemId(id string) *FTextView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FTextView); ok {
			return bt
		}
	}
	return nil
}
func (v *FTextView) setId(s string) *FTextView {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FTextView) setItemId(parent *FListView, id string) *FTextView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}

func getTextViewById(id string) *FTextView {
	if v, ok := GlobalVars.idMap[id].(*FTextView); ok {
		return v
	}
	return nil
}

func (v *FTextView) size(w, h int) *FTextView {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}

func (v *FTextView) background(s string) *FTextView {
	v.FBaseView.background(s)
	return v
}
func (v *FTextView) backgroundColor(s int) *FTextView {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FTextView) cachedBackground(s string) *FTextView {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FTextView) onClick(f func()) *FTextView {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FTextView) visible() *FTextView {
	v.FBaseView.visible()
	return v
}
func (v *FTextView) invisible() *FTextView {
	v.FBaseView.invisible()
	return v
}
func (v *FTextView) gone() *FTextView {
	v.FBaseView.gone()
	return v
}

func (v *FTextView) padding(left, top, right, bottom int) *FTextView {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FTextView) paddingLeft(dp int) *FTextView {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FTextView) paddingTop(dp int) *FTextView {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FTextView) paddingRight(dp int) *FTextView {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FTextView) paddingBottom(dp int) *FTextView {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FTextView) paddingAll(all int) *FTextView {
	v.FBaseView.padding(all, all, all, all)
	return v
}

func (v *FTextView) margin(left, top, right, bottom int) *FTextView {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FTextView) marginLeft(dp int) *FTextView {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FTextView) marginTop(dp int) *FTextView {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FTextView) marginRight(dp int) *FTextView {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FTextView) marginBottom(dp int) *FTextView {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FTextView) marginAll(dp int) *FTextView {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}
func (v *FTextView) layoutGravity(gravity int) *FTextView {
	v.FBaseView.layoutGravity(gravity)
	return v
}

func (v *FTextView) elevation(dp float32) *FTextView {
	v.FBaseView.elevation(dp)
	return v
}
func (v *FTextView) assign(fb **FTextView) *FTextView {
	*fb = v
	return v
}
func (v *FTextView) layoutWeight(f int) *FTextView {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FTextView) text(s string) *FTextView {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}

func (v *FTextView) textColor(s string) *FTextView {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}

func (v *FTextView) textSize(dpsize int) *FTextView {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextSize", sPrintf(dpsize))
	return v
}
func (v *FTextView) getText() string {
	return GlobalVars.uis[v.UI].ViewGetAttr(v.VID, "Text")
}
