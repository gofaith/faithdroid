package faithdroid

type FFrameLayout struct {
	FBaseView
	showAfter bool
	Children  []IBaseView
}

func (vh *ViewHolder) GetFrameLayoutByItemId(id string) *FFrameLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.BaseMap[v].(*FFrameLayout); ok {
			return bt
		}
	}
	return nil
}
func (v *FFrameLayout) SetId(s string) *FFrameLayout {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FFrameLayout) SetItemId(parent *FListView, id string) *FFrameLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetFrameLayoutById(id string) *FFrameLayout {
	if v, ok := GlobalVars.IdMap[id].(*FFrameLayout); ok {
		return v
	}
	return nil
}
func FrameLayout(a IActivity) *FFrameLayout {
	v := &FFrameLayout{}
	v.VID = NewToken()
	v.ClassName = "FrameLayout"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.BaseMap[v.VID] = v
	return v
}
func (v *FFrameLayout) Size(w, h int) *FFrameLayout {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FFrameLayout) X(x float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FFrameLayout) Y(y float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FFrameLayout) PivotX(x float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FFrameLayout) PivotY(y float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FFrameLayout) ScaleX(x float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FFrameLayout) ScaleY(y float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FFrameLayout) Rotation(r float64) *FFrameLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FFrameLayout) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FFrameLayout) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FFrameLayout) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FFrameLayout) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FFrameLayout) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FFrameLayout) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FFrameLayout) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FFrameLayout) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FFrameLayout) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FFrameLayout) Background(s string) *FFrameLayout {
	v.FBaseView.Background(s)
	return v
}
func (v *FFrameLayout) Foreground(s string) *FFrameLayout {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FFrameLayout) CachedForeground(s string) *FFrameLayout {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FFrameLayout) BackgroundColor(s string) *FFrameLayout {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FFrameLayout) CachedBackground(s string) *FFrameLayout {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FFrameLayout) onClick(f func()) *FFrameLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FFrameLayout) Visible() *FFrameLayout {
	v.FBaseView.Visible()
	return v
}
func (v *FFrameLayout) Invisible() *FFrameLayout {
	v.FBaseView.Invisible()
	return v
}
func (v *FFrameLayout) Gone() *FFrameLayout {
	v.FBaseView.Gone()
	return v
}

func (v *FFrameLayout) Padding(left, top, right, bottom int) *FFrameLayout {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FFrameLayout) PaddingLeft(dp int) *FFrameLayout {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FFrameLayout) PaddingTop(dp int) *FFrameLayout {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FFrameLayout) PaddingRight(dp int) *FFrameLayout {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FFrameLayout) PaddingBottom(dp int) *FFrameLayout {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FFrameLayout) PaddingAll(all int) *FFrameLayout {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FFrameLayout) Margin(left, top, right, bottom int) *FFrameLayout {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FFrameLayout) MarginLeft(dp int) *FFrameLayout {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FFrameLayout) MarginTop(dp int) *FFrameLayout {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FFrameLayout) MarginRight(dp int) *FFrameLayout {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FFrameLayout) MarginBottom(dp int) *FFrameLayout {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FFrameLayout) MarginAll(dp int) *FFrameLayout {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FFrameLayout) LayoutGravity(gravity int) *FFrameLayout {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FFrameLayout) Elevation(dp float64) *FFrameLayout {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FFrameLayout) Assign(fb **FFrameLayout) *FFrameLayout {
	*fb = v
	return v
}
func (v *FFrameLayout) LayoutWeight(f int) *FFrameLayout {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FFrameLayout) OnTouch(f func(TouchEvent)) *FFrameLayout {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FFrameLayout) OnClick(f func()) *FFrameLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FFrameLayout) TopToTop(iv IBaseView) *FFrameLayout {
	v.FBaseView.TopToTop(iv)
	return v
}
func (v *FFrameLayout) TopToBottom(iv IBaseView) *FFrameLayout {
	v.FBaseView.TopToBottom(iv)
	return v
}
func (v *FFrameLayout) BottomToBottom(iv IBaseView) *FFrameLayout {
	v.FBaseView.BottomToBottom(iv)
	return v
}
func (v *FFrameLayout) BottomToTop(iv IBaseView) *FFrameLayout {
	v.FBaseView.BottomToTop(iv)
	return v
}
func (v *FFrameLayout) LeftToLeft(iv IBaseView) *FFrameLayout {
	v.FBaseView.LeftToLeft(iv)
	return v
}
func (v *FFrameLayout) LeftToRight(iv IBaseView) *FFrameLayout {
	v.FBaseView.LeftToRight(iv)
	return v
}
func (v *FFrameLayout) RightToRight(iv IBaseView) *FFrameLayout {
	v.FBaseView.RightToRight(iv)
	return v
}
func (v *FFrameLayout) RightToLeft(iv IBaseView) *FFrameLayout {
	v.FBaseView.RightToLeft(iv)
	return v
}
func (v *FFrameLayout) GetBaseView() *FBaseView {
	return &v.FBaseView
}
// --------------------------------------------------------
func (v *FFrameLayout) Append(vs ...IBaseView) *FFrameLayout {
	v.Children = vs
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FFrameLayout) DeferShow() *FFrameLayout {
	v.showAfter = true
	return v
}
