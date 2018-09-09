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
func (v *FViewPager) SetId(s string) *FViewPager {
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

func (v *FViewPager) X(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FViewPager) Y(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FViewPager) PivotX(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FViewPager) PivotY(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FViewPager) ScaleX(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FViewPager) ScaleY(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FViewPager) Rotation(r float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FViewPager) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FViewPager) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FViewPager) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FViewPager) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FViewPager) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FViewPager) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FViewPager) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FViewPager) Background(s string) *FViewPager {
	v.FBaseView.Background(s)
	return v
}
func (v *FViewPager) BackgroundColor(s string) *FViewPager {
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
func (v *FViewPager) Elevation(dp float64) *FViewPager {
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
func (v *FViewPager) OnGetPage(getView func(pos int) IView, getCount func() int) *FViewPager {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(s string) string {
		i, e := a2i(s)
		if e != nil {
			print("OnCreateView():", e.Error())
			return ""
		}
		return getView(i).GetViewId()
	}
	fnId2 := NewToken()
	GlobalVars.EventHandlersMap[fnId2] = func(string) string {
		return SPrintf(getCount())
	}

	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnCreateView", fnId)
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId2)
	return v
}
func (v *FViewPager) BindTabLayoutById(id string) *FViewPager {
	if iview, ok := GlobalVars.IdMap[id]; ok {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TabLayout", iview.GetViewId())
	}
	return v
}
