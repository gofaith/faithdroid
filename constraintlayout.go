package faithdroid

type FConstraintLayout struct {
	FBaseView
	showAfter bool
	Children  []IBaseView
}

var Parent IBaseView = nil

func (vh *ViewHolder) GetConstraintLayoutByItemId(id string) *FConstraintLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.BaseMap[v].(*FConstraintLayout); ok {
			return bt
		}
	}
	return nil
}
func (v *FConstraintLayout) SetId(s string) *FConstraintLayout {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FConstraintLayout) SetItemId(parent *FListView, id string) *FConstraintLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetConstraintLayoutById(id string) *FConstraintLayout {
	if v, ok := GlobalVars.IdMap[id].(*FConstraintLayout); ok {
		return v
	}
	return nil
}
func ConstraintLayout(a *Activity) *FConstraintLayout {
	v := &FConstraintLayout{}
	v.VID = NewToken()
	v.ClassName = "ConstraintLayout"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.BaseMap[v.VID] = v
	return v
}
func (v *FConstraintLayout) Size(w, h int) *FConstraintLayout {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FConstraintLayout) X(x float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FConstraintLayout) Y(y float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FConstraintLayout) PivotX(x float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FConstraintLayout) PivotY(y float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FConstraintLayout) ScaleX(x float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FConstraintLayout) ScaleY(y float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FConstraintLayout) Rotation(r float64) *FConstraintLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FConstraintLayout) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FConstraintLayout) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FConstraintLayout) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FConstraintLayout) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FConstraintLayout) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FConstraintLayout) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FConstraintLayout) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FConstraintLayout) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FConstraintLayout) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FConstraintLayout) Background(s string) *FConstraintLayout {
	v.FBaseView.Background(s)
	return v
}
func (v *FConstraintLayout) Foreground(s string) *FConstraintLayout {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FConstraintLayout) CachedForeground(s string) *FConstraintLayout {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FConstraintLayout) BackgroundColor(s string) *FConstraintLayout {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FConstraintLayout) CachedBackground(s string) *FConstraintLayout {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FConstraintLayout) onClick(f func()) *FConstraintLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FConstraintLayout) Visible() *FConstraintLayout {
	v.FBaseView.Visible()
	return v
}
func (v *FConstraintLayout) Invisible() *FConstraintLayout {
	v.FBaseView.Invisible()
	return v
}
func (v *FConstraintLayout) Gone() *FConstraintLayout {
	v.FBaseView.Gone()
	return v
}

func (v *FConstraintLayout) Padding(left, top, right, bottom int) *FConstraintLayout {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FConstraintLayout) PaddingLeft(dp int) *FConstraintLayout {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FConstraintLayout) PaddingTop(dp int) *FConstraintLayout {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FConstraintLayout) PaddingRight(dp int) *FConstraintLayout {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FConstraintLayout) PaddingBottom(dp int) *FConstraintLayout {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FConstraintLayout) PaddingAll(all int) *FConstraintLayout {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FConstraintLayout) Margin(left, top, right, bottom int) *FConstraintLayout {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FConstraintLayout) MarginLeft(dp int) *FConstraintLayout {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FConstraintLayout) MarginTop(dp int) *FConstraintLayout {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FConstraintLayout) MarginRight(dp int) *FConstraintLayout {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FConstraintLayout) MarginBottom(dp int) *FConstraintLayout {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FConstraintLayout) MarginAll(dp int) *FConstraintLayout {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FConstraintLayout) LayoutGravity(gravity int) *FConstraintLayout {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FConstraintLayout) Elevation(dp float64) *FConstraintLayout {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FConstraintLayout) Assign(fb **FConstraintLayout) *FConstraintLayout {
	*fb = v
	return v
}
func (v *FConstraintLayout) LayoutWeight(f int) *FConstraintLayout {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FConstraintLayout) OnTouch(f func(TouchEvent)) *FConstraintLayout {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FConstraintLayout) OnClick(f func()) *FConstraintLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FConstraintLayout) TopToTop(iv IBaseView) *FConstraintLayout {
	v.FBaseView.TopToTop(iv)
	return v
}
func (v *FConstraintLayout) TopToBottom(iv IBaseView) *FConstraintLayout {
	v.FBaseView.TopToBottom(iv)
	return v
}
func (v *FConstraintLayout) BottomToBottom(iv IBaseView) *FConstraintLayout {
	v.FBaseView.BottomToBottom(iv)
	return v
}
func (v *FConstraintLayout) BottomToTop(iv IBaseView) *FConstraintLayout {
	v.FBaseView.BottomToTop(iv)
	return v
}
func (v *FConstraintLayout) LeftToLeft(iv IBaseView) *FConstraintLayout {
	v.FBaseView.LeftToLeft(iv)
	return v
}
func (v *FConstraintLayout) LeftToRight(iv IBaseView) *FConstraintLayout {
	v.FBaseView.LeftToRight(iv)
	return v
}
func (v *FConstraintLayout) RightToRight(iv IBaseView) *FConstraintLayout {
	v.FBaseView.RightToRight(iv)
	return v
}
func (v *FConstraintLayout) RightToLeft(iv IBaseView) *FConstraintLayout {
	v.FBaseView.RightToLeft(iv)
	return v
}
func (v *FConstraintLayout) GetBaseView() *FBaseView {
	return &v.FBaseView
}

// --------------------------------------------------------
func (v *FConstraintLayout) Append(vs ...IBaseView) *FConstraintLayout {
	v.Children = append(v.Children, vs...)
	for _, i := range vs {
		if i.GetBaseView().Fn_TopToTop != nil {
			i.GetBaseView().Fn_TopToTop()
		}
		if i.GetBaseView().Fn_TopToBottom != nil {
			i.GetBaseView().Fn_TopToBottom()
		}
		if i.GetBaseView().Fn_BottomToBottom != nil {
			i.GetBaseView().Fn_BottomToBottom()
		}
		if i.GetBaseView().Fn_BottomToTop != nil {
			i.GetBaseView().Fn_BottomToTop()
		}
		if i.GetBaseView().Fn_LeftToLeft != nil {
			i.GetBaseView().Fn_LeftToLeft()
		}
		if i.GetBaseView().Fn_LeftToRight != nil {
			i.GetBaseView().Fn_LeftToRight()
		}
		if i.GetBaseView().Fn_RightToRight != nil {
			i.GetBaseView().Fn_RightToRight()
		}
		if i.GetBaseView().Fn_RightToLeft != nil {
			i.GetBaseView().Fn_RightToLeft()
		}
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FConstraintLayout) DeferShow() *FConstraintLayout {
	v.showAfter = true
	return v
}
