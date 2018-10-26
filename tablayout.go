package faithdroid

type FTabLayout struct {
	FBaseView
}

func (vh *ViewHolder) GetTabLayoutByItemId(id string) *FTabLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FTabLayout); ok {
			return bt
		}
	}
	return nil
}

func (v *FTabLayout) SetId(s string) *FTabLayout {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FTabLayout) SetItemId(parent *FListView, id string) *FTabLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetTabLayoutById(id string) *FTabLayout {
	if v, ok := GlobalVars.IdMap[id].(*FTabLayout); ok {
		return v
	}
	return nil
}
func TabLayout(a IActivity) *FTabLayout {
	v := &FTabLayout{}
	v.VID = NewToken()
	v.ClassName = "TabLayout"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FTabLayout) Size(w, h int) *FTabLayout {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FTabLayout) X(x float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FTabLayout) Y(y float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FTabLayout) PivotX(x float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FTabLayout) PivotY(y float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FTabLayout) ScaleX(x float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FTabLayout) ScaleY(y float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FTabLayout) Rotation(r float64) *FTabLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FTabLayout) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FTabLayout) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FTabLayout) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FTabLayout) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FTabLayout) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FTabLayout) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FTabLayout) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FTabLayout) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FTabLayout) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FTabLayout) Background(s string) *FTabLayout {
	v.FBaseView.Background(s)
	return v
}
func (v *FTabLayout) Foreground(s string) *FTabLayout {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FTabLayout) CachedForeground(s string) *FTabLayout {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FTabLayout) BackgroundColor(s string) *FTabLayout {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FTabLayout) CachedBackground(s string) *FTabLayout {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FTabLayout) onClick(f func()) *FTabLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FTabLayout) Visible() *FTabLayout {
	v.FBaseView.Visible()
	return v
}
func (v *FTabLayout) Invisible() *FTabLayout {
	v.FBaseView.Invisible()
	return v
}
func (v *FTabLayout) Gone() *FTabLayout {
	v.FBaseView.Gone()
	return v
}

func (v *FTabLayout) Padding(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FTabLayout) PaddingLeft(dp int) *FTabLayout {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) PaddingTop(dp int) *FTabLayout {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) PaddingRight(dp int) *FTabLayout {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) PaddingBottom(dp int) *FTabLayout {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) PaddingAll(all int) *FTabLayout {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FTabLayout) Margin(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FTabLayout) MarginLeft(dp int) *FTabLayout {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) MarginTop(dp int) *FTabLayout {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) MarginRight(dp int) *FTabLayout {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) MarginBottom(dp int) *FTabLayout {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) MarginAll(dp int) *FTabLayout {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FTabLayout) LayoutGravity(gravity int) *FTabLayout {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FTabLayout) Elevation(dp float64) *FTabLayout {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FTabLayout) Assign(fb **FTabLayout) *FTabLayout {
	*fb = v
	return v
}
func (v *FTabLayout) LayoutWeight(f int) *FTabLayout {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FTabLayout) OnTouch(f func(TouchEvent)) *FTabLayout {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FTabLayout) OnClick(f func()) *FTabLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FTabLayout) Clickable(b bool) *FTabLayout {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FTabLayout) Top2TopOf(id string) *FTabLayout {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FTabLayout) Top2BottomOf(id string) *FTabLayout {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FTabLayout) Bottom2TopOf(id string) *FTabLayout {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FTabLayout) Bottom2BottomOf(id string) *FTabLayout {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FTabLayout) Left2LeftOf(id string) *FTabLayout {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FTabLayout) Right2RightOf(id string) *FTabLayout {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FTabLayout) Left2RightOf(id string) *FTabLayout {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FTabLayout) Right2LeftOf(id string) *FTabLayout {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FTabLayout) CenterX() *FTabLayout {
	v.FBaseView.CenterX()
	return v
}
func (v *FTabLayout) CenterY() *FTabLayout {
	v.FBaseView.CenterY()
	return v
}
// ------------------------------------------------------------

type FTab struct {
	Text, Icon string
}

func Tab(text string) *FTab {
	t := &FTab{}
	t.Text = text
	return t
}
func (t *FTab) IconSrc(s string) *FTab {
	t.Icon = s
	return t
}

// ----------------------------------------------------------

func (v *FTabLayout) Tabs(ts ...*FTab) *FTabLayout {
	for _, t := range ts {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddTab", JsonObject(t))
	}
	return v
}
