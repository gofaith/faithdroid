package faithdroid

type FProgressBar struct {
	FBaseView
}

func ProgressBar(a IActivity) *FProgressBar {
	v := &FProgressBar{}
	v.VID = NewToken()
	v.ClassName = "ProgressBar"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetProgressBarByItemId(id string) *FProgressBar {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FProgressBar); ok {
			return bt
		}
	}
	return nil
}

func (v *FProgressBar) Size(w, h int) *FProgressBar {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FProgressBar) X(x float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FProgressBar) Y(y float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FProgressBar) PivotX(x float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FProgressBar) PivotY(y float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FProgressBar) ScaleX(x float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FProgressBar) ScaleY(y float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FProgressBar) Rotation(r float64) *FProgressBar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FProgressBar) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FProgressBar) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FProgressBar) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FProgressBar) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FProgressBar) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FProgressBar) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FProgressBar) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FProgressBar) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FProgressBar) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FProgressBar) SetId(s string) *FProgressBar {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FProgressBar) SetItemId(parent *FListView, id string) *FProgressBar {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetProgressBarById(id string) *FProgressBar {
	if v, ok := GlobalVars.IdMap[id].(*FProgressBar); ok {
		return v
	}
	return nil
}

func (v *FProgressBar) Background(s string) *FProgressBar {
	v.FBaseView.Background(s)
	return v
}
func (v *FProgressBar) Foreground(s string) *FProgressBar {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FProgressBar) CachedForeground(s string) *FProgressBar {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FProgressBar) BackgroundColor(s string) *FProgressBar {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FProgressBar) CachedBackground(s string) *FProgressBar {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FProgressBar) Visible() *FProgressBar {
	v.FBaseView.Visible()
	return v
}
func (v *FProgressBar) Invisible() *FProgressBar {
	v.FBaseView.Invisible()
	return v
}
func (v *FProgressBar) Gone() *FProgressBar {
	v.FBaseView.Gone()
	return v
}

func (v *FProgressBar) Padding(left, top, right, bottom int) *FProgressBar {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FProgressBar) PaddingLeft(dp int) *FProgressBar {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FProgressBar) PaddingTop(dp int) *FProgressBar {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FProgressBar) PaddingRight(dp int) *FProgressBar {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FProgressBar) PaddingBottom(dp int) *FProgressBar {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FProgressBar) PaddingAll(dp int) *FProgressBar {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FProgressBar) Margin(left, top, right, bottom int) *FProgressBar {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FProgressBar) MarginLeft(dp int) *FProgressBar {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FProgressBar) MarginTop(dp int) *FProgressBar {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FProgressBar) MarginRight(dp int) *FProgressBar {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FProgressBar) MarginBottom(dp int) *FProgressBar {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FProgressBar) MarginAll(dp int) *FProgressBar {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FProgressBar) LayoutGravity(gravity int) *FProgressBar {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FProgressBar) Elevation(dp float64) *FProgressBar {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FProgressBar) Assign(fb **FProgressBar) *FProgressBar {
	*fb = v
	return v
}
func (v *FProgressBar) LayoutWeight(f int) *FProgressBar {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FProgressBar) OnTouch(f func(TouchEvent)) *FProgressBar {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FProgressBar) OnClick(f func()) *FProgressBar {
	v.FBaseView.OnClick(f)
	return v
}
func (v *FProgressBar) Clickable(b bool) *FProgressBar {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FProgressBar) Top2TopOf(id string) *FProgressBar {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FProgressBar) Top2BottomOf(id string) *FProgressBar {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FProgressBar) Bottom2TopOf(id string) *FProgressBar {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FProgressBar) Bottom2BottomOf(id string) *FProgressBar {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FProgressBar) Left2LeftOf(id string) *FProgressBar {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FProgressBar) Right2RightOf(id string) *FProgressBar {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FProgressBar) Left2RightOf(id string) *FProgressBar {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FProgressBar) Right2LeftOf(id string) *FProgressBar {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FProgressBar) CenterX() *FProgressBar {
	v.FBaseView.CenterX()
	return v
}
func (v *FProgressBar) CenterY() *FProgressBar {
	v.FBaseView.CenterY()
	return v
}
func (v *FProgressBar) WidthPercent(num float64) *FProgressBar {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FProgressBar) HeightPercent(num float64) *FProgressBar {
	v.FBaseView.HeightPercent(num)
	return v
}

// --------------------------------------------------------
