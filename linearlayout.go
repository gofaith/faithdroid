package faithdroid

type FLinearLayout struct {
	FBaseView
	Children  []IView
	showAfter bool
}

func (vh *ViewHolder) GetLinearLayoutByItemId(id string) *FLinearLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FLinearLayout); ok {
			return bt
		}
	}
	return nil
}
func (v *FLinearLayout) SetId(s string) *FLinearLayout {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FLinearLayout) SetItemId(parent *FListView, id string) *FLinearLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetLinearLayoutById(id string) *FLinearLayout {
	if v, ok := GlobalVars.IdMap[id].(*FLinearLayout); ok {
		return v
	}
	return nil
}
func LinearLayout(a IActivity) *FLinearLayout {
	v := &FLinearLayout{}
	v.VID = NewToken()
	v.ClassName = "LinearLayout"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FLinearLayout) Size(w, h int) *FLinearLayout {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FLinearLayout) X(x float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FLinearLayout) Y(y float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FLinearLayout) PivotX(x float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FLinearLayout) PivotY(y float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FLinearLayout) ScaleX(x float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FLinearLayout) ScaleY(y float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FLinearLayout) Rotation(r float64) *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FLinearLayout) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FLinearLayout) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FLinearLayout) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FLinearLayout) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FLinearLayout) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FLinearLayout) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FLinearLayout) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FLinearLayout) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FLinearLayout) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FLinearLayout) Background(s string) *FLinearLayout {
	v.FBaseView.Background(s)
	return v
}
func (v *FLinearLayout) Foreground(s string) *FLinearLayout {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FLinearLayout) CachedForeground(s string) *FLinearLayout {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FLinearLayout) BackgroundColor(s string) *FLinearLayout {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FLinearLayout) CachedBackground(s string) *FLinearLayout {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FLinearLayout) onClick(f func()) *FLinearLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FLinearLayout) Visible() *FLinearLayout {
	v.FBaseView.Visible()
	return v
}
func (v *FLinearLayout) Invisible() *FLinearLayout {
	v.FBaseView.Invisible()
	return v
}
func (v *FLinearLayout) Gone() *FLinearLayout {
	v.FBaseView.Gone()
	return v
}

func (v *FLinearLayout) Padding(left, top, right, bottom int) *FLinearLayout {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FLinearLayout) PaddingLeft(dp int) *FLinearLayout {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FLinearLayout) PaddingTop(dp int) *FLinearLayout {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FLinearLayout) PaddingRight(dp int) *FLinearLayout {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FLinearLayout) PaddingBottom(dp int) *FLinearLayout {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FLinearLayout) PaddingAll(all int) *FLinearLayout {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FLinearLayout) Margin(left, top, right, bottom int) *FLinearLayout {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FLinearLayout) MarginLeft(dp int) *FLinearLayout {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FLinearLayout) MarginTop(dp int) *FLinearLayout {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FLinearLayout) MarginRight(dp int) *FLinearLayout {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FLinearLayout) MarginBottom(dp int) *FLinearLayout {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FLinearLayout) MarginAll(dp int) *FLinearLayout {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FLinearLayout) LayoutGravity(gravity int) *FLinearLayout {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FLinearLayout) Elevation(dp float64) *FLinearLayout {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FLinearLayout) Assign(fb **FLinearLayout) *FLinearLayout {
	*fb = v
	return v
}
func (v *FLinearLayout) LayoutWeight(f int) *FLinearLayout {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FLinearLayout) OnTouch(f func(TouchEvent)) *FLinearLayout {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FLinearLayout) OnClick(f func()) *FLinearLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FLinearLayout) Clickable(b bool) *FLinearLayout {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FLinearLayout) Top2TopOf(id string) *FLinearLayout {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FLinearLayout) Top2BottomOf(id string) *FLinearLayout {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FLinearLayout) Bottom2TopOf(id string) *FLinearLayout {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FLinearLayout) Bottom2BottomOf(id string) *FLinearLayout {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FLinearLayout) Left2LeftOf(id string) *FLinearLayout {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FLinearLayout) Right2RightOf(id string) *FLinearLayout {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FLinearLayout) Left2RightOf(id string) *FLinearLayout {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FLinearLayout) Right2LeftOf(id string) *FLinearLayout {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FLinearLayout) CenterX() *FLinearLayout {
	v.FBaseView.CenterX()
	return v
}
func (v *FLinearLayout) CenterY() *FLinearLayout {
	v.FBaseView.CenterY()
	return v
}
func (v *FLinearLayout) WidthPercent(num float64) *FLinearLayout {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FLinearLayout) HeightPercent(num float64) *FLinearLayout {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
func (v *FLinearLayout) Append(vs ...IView) *FLinearLayout {
	v.Children = append(v.Children, vs...)
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.showAfter = false
		v.Show()
	}
	return v
}
func (v *FLinearLayout) AddView(i IView) *FLinearLayout {
	v.Append(i)
	return v
}
func (v *FLinearLayout) AddViewAt(i IView, pos int) *FLinearLayout {
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
func (v *FLinearLayout) Vertical() *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Orientation", "VERTICAL")
	return v
}
func (v *FLinearLayout) Horizontal() *FLinearLayout {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Orientation", "HORIZONTAL")
	return v
}
func (v *FLinearLayout) IsVertical() bool {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Orientation") == "VERTICAL"
}
func (v *FLinearLayout) DeferShow() *FLinearLayout {
	v.showAfter = true
	return v
}
