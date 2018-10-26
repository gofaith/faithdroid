package faithdroid

type FHScrollView struct {
	FBaseView
	Children  []IView
	showAfter bool
}

func (vh *ViewHolder) GetHScrollViewByItemId(id string) *FHScrollView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FHScrollView); ok {
			return bt
		}
	}
	return nil
}
func (v *FHScrollView) SetId(s string) *FHScrollView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FHScrollView) SetItemId(parent *FListView, id string) *FHScrollView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetHScrollViewById(id string) *FHScrollView {
	if v, ok := GlobalVars.IdMap[id].(*FHScrollView); ok {
		return v
	}
	return nil
}
func HScrollView(a IActivity) *FHScrollView {
	v := &FHScrollView{}
	v.VID = NewToken()
	v.ClassName = "HScrollView"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FHScrollView) Size(w, h int) *FHScrollView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FHScrollView) X(x float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FHScrollView) Y(y float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FHScrollView) PivotX(x float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FHScrollView) PivotY(y float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FHScrollView) ScaleX(x float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FHScrollView) ScaleY(y float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FHScrollView) Rotation(r float64) *FHScrollView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FHScrollView) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FHScrollView) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FHScrollView) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FHScrollView) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FHScrollView) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FHScrollView) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FHScrollView) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FHScrollView) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FHScrollView) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FHScrollView) Background(s string) *FHScrollView {
	v.FBaseView.Background(s)
	return v
}
func (v *FHScrollView) Foreground(s string) *FHScrollView {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FHScrollView) CachedForeground(s string) *FHScrollView {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FHScrollView) BackgroundColor(s string) *FHScrollView {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FHScrollView) CachedBackground(s string) *FHScrollView {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FHScrollView) onClick(f func()) *FHScrollView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FHScrollView) Visible() *FHScrollView {
	v.FBaseView.Visible()
	return v
}
func (v *FHScrollView) Invisible() *FHScrollView {
	v.FBaseView.Invisible()
	return v
}
func (v *FHScrollView) Gone() *FHScrollView {
	v.FBaseView.Gone()
	return v
}

func (v *FHScrollView) Padding(left, top, right, bottom int) *FHScrollView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FHScrollView) PaddingLeft(dp int) *FHScrollView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FHScrollView) PaddingTop(dp int) *FHScrollView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FHScrollView) PaddingRight(dp int) *FHScrollView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FHScrollView) PaddingBottom(dp int) *FHScrollView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FHScrollView) PaddingAll(all int) *FHScrollView {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FHScrollView) Margin(left, top, right, bottom int) *FHScrollView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FHScrollView) MarginLeft(dp int) *FHScrollView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FHScrollView) MarginTop(dp int) *FHScrollView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FHScrollView) MarginRight(dp int) *FHScrollView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FHScrollView) MarginBottom(dp int) *FHScrollView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FHScrollView) MarginAll(dp int) *FHScrollView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FHScrollView) LayoutGravity(gravity int) *FHScrollView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FHScrollView) Elevation(dp float64) *FHScrollView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FHScrollView) Assign(fb **FHScrollView) *FHScrollView {
	*fb = v
	return v
}
func (v *FHScrollView) LayoutWeight(f int) *FHScrollView {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FHScrollView) OnTouch(f func(TouchEvent)) *FHScrollView {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FHScrollView) OnClick(f func()) *FHScrollView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FHScrollView) Clickable(b bool) *FHScrollView {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FHScrollView) Top2TopOf(id string) *FHScrollView {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FHScrollView) Top2BottomOf(id string) *FHScrollView {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FHScrollView) Bottom2TopOf(id string) *FHScrollView {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FHScrollView) Bottom2BottomOf(id string) *FHScrollView {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FHScrollView) Left2LeftOf(id string) *FHScrollView {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FHScrollView) Right2RightOf(id string) *FHScrollView {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FHScrollView) Left2RightOf(id string) *FHScrollView {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FHScrollView) Right2LeftOf(id string) *FHScrollView {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FHScrollView) CenterX() *FHScrollView {
	v.FBaseView.CenterX()
	return v
}
func (v *FHScrollView) CenterY() *FHScrollView {
	v.FBaseView.CenterY()
	return v
}
func (v *FHScrollView) WidthPercent(num float64) *FHScrollView {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FHScrollView) HeightPercent(num float64) *FHScrollView {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
func (v *FHScrollView) Append(vs ...IView) *FHScrollView {
	v.Children = vs
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.showAfter = false
		v.Show()
	}
	return v
}

func (v *FHScrollView) AddView(i IView) *FHScrollView {
	v.Append(i)
	return v
}
func (v *FHScrollView) AddViewAt(i IView, pos int) *FHScrollView {
	if pos == -1 {
		v.AddView(i)
		return v
	}
	if pos < 0 {
		pos = len(v.Children) + pos + 1
	}
	v.Children = append(v.Children[:pos], append([]IView{i}, v.Children[pos:]...)...)
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddViewAt", JsonArray([]string{SPrintf(pos), i.GetViewId()}))
	return v
}
func (v *FHScrollView) DeferShow() *FHScrollView {
	v.showAfter = true
	return v
}
