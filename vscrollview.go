package faithdroid

type FVScrollView struct {
	FBaseView
	showAfter bool
	Children  []IView
}

func (vh *ViewHolder) GetVScrollViewByItemId(id string) *FVScrollView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FVScrollView); ok {
			return bt
		}
	}
	return nil
}
func (v *FVScrollView) SetId(s string) *FVScrollView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FVScrollView) SetItemId(parent *FListView, id string) *FVScrollView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetVScrollViewById(id string) *FVScrollView {
	if v, ok := GlobalVars.IdMap[id].(*FVScrollView); ok {
		return v
	}
	return nil
}
func VScrollView(a *Activity) *FVScrollView {
	v := &FVScrollView{}
	v.VID = NewToken()
	v.ClassName = "VScrollView"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FVScrollView) Size(w, h int) *FVScrollView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FVScrollView) X(x float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FVScrollView) Y(y float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FVScrollView) PivotX(x float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FVScrollView) PivotY(y float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FVScrollView) ScaleX(x float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FVScrollView) ScaleY(y float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FVScrollView) Rotation(r float64) *FVScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FVScrollView) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FVScrollView) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FVScrollView) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FVScrollView) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FVScrollView) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FVScrollView) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FVScrollView) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FVScrollView) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FVScrollView) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FVScrollView) Background(s string) *FVScrollView {
	v.FBaseView.Background(s)
	return v
}
func (v *FVScrollView) Foreground(s string) *FVScrollView {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FVScrollView) CachedForeground(s string) *FVScrollView {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FVScrollView) BackgroundColor(s string) *FVScrollView {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FVScrollView) CachedBackground(s string) *FVScrollView {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FVScrollView) onClick(f func()) *FVScrollView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FVScrollView) Visible() *FVScrollView {
	v.FBaseView.Visible()
	return v
}
func (v *FVScrollView) Invisible() *FVScrollView {
	v.FBaseView.Invisible()
	return v
}
func (v *FVScrollView) Gone() *FVScrollView {
	v.FBaseView.Gone()
	return v
}

func (v *FVScrollView) Padding(left, top, right, bottom int) *FVScrollView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FVScrollView) PaddingLeft(dp int) *FVScrollView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FVScrollView) PaddingTop(dp int) *FVScrollView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FVScrollView) PaddingRight(dp int) *FVScrollView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FVScrollView) PaddingBottom(dp int) *FVScrollView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FVScrollView) PaddingAll(all int) *FVScrollView {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FVScrollView) Margin(left, top, right, bottom int) *FVScrollView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FVScrollView) MarginLeft(dp int) *FVScrollView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FVScrollView) MarginTop(dp int) *FVScrollView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FVScrollView) MarginRight(dp int) *FVScrollView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FVScrollView) MarginBottom(dp int) *FVScrollView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FVScrollView) MarginAll(dp int) *FVScrollView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FVScrollView) LayoutGravity(gravity int) *FVScrollView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FVScrollView) Elevation(dp float64) *FVScrollView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FVScrollView) Assign(fb **FVScrollView) *FVScrollView {
	*fb = v
	return v
}
func (v *FVScrollView) LayoutWeight(f int) *FVScrollView {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FVScrollView) OnTouch(f func(TouchEvent)) *FVScrollView {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FVScrollView) OnClick(f func()) *FVScrollView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

// --------------------------------------------------------
func (v *FVScrollView) Append(vs ...IView) *FVScrollView {
	v.Children = vs
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FVScrollView) DeferShow() *FVScrollView {
	v.showAfter = true
	return v
}
