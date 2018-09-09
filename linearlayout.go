package faithdroid

type FLinearLayout struct {
	FBaseView
	showAfter bool
	Children  []IView
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
func LinearLayout(a *Activity) *FLinearLayout {
	v := &FLinearLayout{}
	v.VID = NewToken()
	v.ClassName = "LinearLayout"
	v.UI = a.UI
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

// --------------------------------------------------------
func (v *FLinearLayout) Append(vs ...IView) *FLinearLayout {
	v.Children = vs
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	if v.showAfter {
		v.Show()
	}
	return v
}
func (v *FLinearLayout) DeferShow() *FLinearLayout {
	v.showAfter = true
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
