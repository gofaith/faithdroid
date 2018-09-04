package faithdroid

type FViewPager struct {
	FBaseView
}

func (vh *ViewHolder) getViewPagerByItemId(id string) *FViewPager {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FViewPager); ok {
			return bt
		}
	}
	return nil
}
func (v *FViewPager) setId(s string) *FViewPager {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FViewPager) setItemId(parent *FListView, id string) *FViewPager {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getViewPagerById(id string) *FViewPager {
	if v, ok := GlobalVars.idMap[id].(*FViewPager); ok {
		return v
	}
	return nil
}
func viewpager(a *Activity) *FViewPager {
	v := &FViewPager{}
	v.VID = newToken()
	v.ClassName = "ViewPager"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (v *FViewPager) size(w, h int) *FViewPager {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}

func (v *FViewPager) background(s string) *FViewPager {
	v.FBaseView.background(s)
	return v
}
func (v *FViewPager) backgroundColor(s int) *FViewPager {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FViewPager) cachedBackground(s string) *FViewPager {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FViewPager) onClick(f func()) *FViewPager {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FViewPager) visible() *FViewPager {
	v.FBaseView.visible()
	return v
}
func (v *FViewPager) invisible() *FViewPager {
	v.FBaseView.invisible()
	return v
}
func (v *FViewPager) gone() *FViewPager {
	v.FBaseView.gone()
	return v
}

func (v *FViewPager) padding(left, top, right, bottom int) *FViewPager {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FViewPager) paddingLeft(dp int) *FViewPager {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) paddingTop(dp int) *FViewPager {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FViewPager) paddingRight(dp int) *FViewPager {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FViewPager) paddingBottom(dp int) *FViewPager {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FViewPager) paddingAll(all int) *FViewPager {
	v.FBaseView.padding(all, all, all, all)
	return v
}

func (v *FViewPager) margin(left, top, right, bottom int) *FViewPager {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FViewPager) marginLeft(dp int) *FViewPager {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) marginTop(dp int) *FViewPager {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FViewPager) marginRight(dp int) *FViewPager {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FViewPager) marginBottom(dp int) *FViewPager {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FViewPager) marginAll(dp int) *FViewPager {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FViewPager) layoutGravity(gravity int) *FViewPager {
	v.FBaseView.layoutGravity(gravity)
	return v
}

func (v *FViewPager) elevation(dp float32) *FViewPager {
	v.FBaseView.elevation(dp)
	return v
}
func (v *FViewPager) assign(fb **FViewPager) *FViewPager {
	*fb = v
	return v
}

func (v *FViewPager) layoutWeight(f int) *FViewPager {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
type FPage struct {
	VID, OnSelected string
}

func page(createView func() IView, onSelected func()) *FPage {
	p := &FPage{}
	p.VID = newToken()
	GlobalVars.eventHandlersMap[p.VID] = func(string) string {
		return createView().getVID()
	}
	if onSelected == nil {
		return p
	}
	p.OnSelected = newToken()
	GlobalVars.eventHandlersMap[p.OnSelected] = func(string) string {
		onSelected()
		return ""
	}
	return p
}

func (v *FViewPager) pages(ps ...*FPage) *FViewPager {
	if ps == nil {
		return v
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Pages", jsonArray(ps))
	return v
}
