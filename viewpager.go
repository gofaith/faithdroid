package faithdroid

type FViewPager struct {
	FBaseView
}

func (vh *ViewHolder) GetViewPagerByItemId(id string) *FViewPager {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FViewPager); ok {
			return bt
		}
	}
	return nil
}
func (v *FViewPager) SetId(s string)*FViewPager {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FViewPager) SetItemId(parent *FListView, id string) *FViewPager {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetViewPagerById(id string) *FViewPager {
	if v, ok := GlobalVars.IdMap[id].(*FViewPager); ok {
		return v
	}
	return nil
}
func ViewPager(a *Activity) *FViewPager {
	v := &FViewPager{}
	v.VID = NewToken()
	v.ClassName = "ViewPager"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FViewPager) Size(w, h int) *FViewPager {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FViewPager) Background(s string) *FViewPager {
	v.FBaseView.Background(s)
	return v
}
func (v *FViewPager) BackgroundColor(s int) *FViewPager {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FViewPager) CachedBackground(s string) *FViewPager {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FViewPager) onClick(f func()) *FViewPager {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FViewPager) Visible() *FViewPager {
	v.FBaseView.Visible()
	return v
}
func (v *FViewPager) Invisible() *FViewPager {
	v.FBaseView.Invisible()
	return v
}
func (v *FViewPager) Gone() *FViewPager {
	v.FBaseView.Gone()
	return v
}

func (v *FViewPager) Padding(left, top, right, bottom int) *FViewPager {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FViewPager) PaddingLeft(dp int) *FViewPager {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) PaddingTop(dp int) *FViewPager {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FViewPager) PaddingRight(dp int) *FViewPager {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FViewPager) PaddingBottom(dp int) *FViewPager {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FViewPager) PaddingAll(all int) *FViewPager {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FViewPager) Margin(left, top, right, bottom int) *FViewPager {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FViewPager) MarginLeft(dp int) *FViewPager {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) MarginTop(dp int) *FViewPager {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FViewPager) MarginRight(dp int) *FViewPager {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FViewPager) MarginBottom(dp int) *FViewPager {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FViewPager) MarginAll(dp int) *FViewPager {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FViewPager) LayoutGravity(gravity int) *FViewPager {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FViewPager) Elevation(dp float32) *FViewPager {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FViewPager) Assign(fb **FViewPager) *FViewPager {
	*fb = v
	return v
}
func (v *FViewPager) LayoutWeight(f int) *FViewPager {
	v.FBaseView.LayoutWeight(f)
	return v
}

// --------------------------------------------------------
type FPage struct {
	VID string
}

func Page(createView func() IView) *FPage {
	p := &FPage{}
	p.VID = NewToken()
	GlobalVars.EventHandlersMap[p.VID] = func(string) string {
		return createView().GetViewId()
	}
	return p
}

func (v *FViewPager) Pages(ps ...*FPage) *FViewPager {
	if ps == nil {
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Pages", JsonArray(ps))
	return v
}
func (v *FViewPager) BindTabLayoutById(id string) *FViewPager {
	if iview, ok := GlobalVars.IdMap[id]; ok {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TabLayout", iview.GetViewId())
	}
	return v
}
