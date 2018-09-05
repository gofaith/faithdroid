package faithdroid

type FTabLayout struct {
	FBaseView
}

func (vh *ViewHolder) getTabLayoutByItemId(id string) *FTabLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FTabLayout); ok {
			return bt
		}
	}
	return nil
}

func (v *FTabLayout) setId(s string) *FTabLayout {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FTabLayout) setItemId(parent *FListView, id string) *FTabLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getTabLayoutById(id string) *FTabLayout {
	if v, ok := GlobalVars.idMap[id].(*FTabLayout); ok {
		return v
	}
	return nil
}
func TabLayout(a *Activity) *FTabLayout {
	v := &FTabLayout{}
	v.VID = newToken()
	v.ClassName = "TabLayout"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (v *FTabLayout) size(w, h int) *FTabLayout {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}

func (v *FTabLayout) background(s string) *FTabLayout {
	v.FBaseView.background(s)
	return v
}
func (v *FTabLayout) backgroundColor(s int) *FTabLayout {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FTabLayout) cachedBackground(s string) *FTabLayout {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FTabLayout) onClick(f func()) *FTabLayout {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FTabLayout) visible() *FTabLayout {
	v.FBaseView.visible()
	return v
}
func (v *FTabLayout) invisible() *FTabLayout {
	v.FBaseView.invisible()
	return v
}
func (v *FTabLayout) gone() *FTabLayout {
	v.FBaseView.gone()
	return v
}

func (v *FTabLayout) padding(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FTabLayout) paddingLeft(dp int) *FTabLayout {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) paddingTop(dp int) *FTabLayout {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) paddingRight(dp int) *FTabLayout {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) paddingBottom(dp int) *FTabLayout {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) paddingAll(all int) *FTabLayout {
	v.FBaseView.padding(all, all, all, all)
	return v
}

func (v *FTabLayout) margin(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FTabLayout) marginLeft(dp int) *FTabLayout {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) marginTop(dp int) *FTabLayout {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) marginRight(dp int) *FTabLayout {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) marginBottom(dp int) *FTabLayout {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) marginAll(dp int) *FTabLayout {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FTabLayout) layoutGravity(gravity int) *FTabLayout {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FTabLayout) elevation(dp float32) *FTabLayout {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FTabLayout) assign(fb **FTabLayout) *FTabLayout {
	*fb = v
	return v
}
func (v *FTabLayout) layoutWeight(f int) *FTabLayout {
	v.FBaseView.layoutWeight(f)
	return v
}

// ------------------------------------------------------------

type FTab struct {
	Text, Icon string
}

func tab(text string, i ...string) *FTab {
	t := &FTab{}
	t.Text = text
	if len(i) > 0 {
		t.Icon = i[0]
	}
	return t
}

// ----------------------------------------------------------

func (v *FTabLayout) tabs(ts ...*FTab) *FTabLayout {
	for _, t := range ts {
		GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "AddTab", jsonObject(t))
	}
	return v
}
