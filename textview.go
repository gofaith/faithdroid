package faithdroid

type FTextView struct {
	FBaseView
}

func TextView(a *Activity) *FTextView {
	v := &FTextView{}
	v.VID = NewToken()
	v.ClassName = "TextView"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetTextViewByItemId(id string) *FTextView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FTextView); ok {
			return bt
		}
	}
	return nil
}
func (v *FTextView) SetId(s string) *FTextView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FTextView) SetItemId(parent *FListView, id string) *FTextView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}

func GetTextViewById(id string) *FTextView {
	if v, ok := GlobalVars.IdMap[id].(*FTextView); ok {
		return v
	}
	return nil
}

func (v *FTextView) Size(w, h int) *FTextView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FTextView) X(x float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FTextView) Y(y float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FTextView) PivotX(x float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FTextView) PivotY(y float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FTextView) ScaleX(x float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FTextView) ScaleY(y float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FTextView) Rotation(r float64) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FTextView) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FTextView) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FTextView) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FTextView) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FTextView) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FTextView) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FTextView) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FTextView) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FTextView) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FTextView) Background(s string) *FTextView {
	v.FBaseView.Background(s)
	return v
}
func (v *FTextView) Foreground(s string) *FTextView {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FTextView) CachedForeground(s string) *FTextView {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FTextView) BackgroundColor(s string) *FTextView {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FTextView) CachedBackground(s string) *FTextView {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FTextView) onClick(f func()) *FTextView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FTextView) Visible() *FTextView {
	v.FBaseView.Visible()
	return v
}
func (v *FTextView) Invisible() *FTextView {
	v.FBaseView.Invisible()
	return v
}
func (v *FTextView) Gone() *FTextView {
	v.FBaseView.Gone()
	return v
}

func (v *FTextView) Padding(left, top, right, bottom int) *FTextView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FTextView) PaddingLeft(dp int) *FTextView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FTextView) PaddingTop(dp int) *FTextView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FTextView) PaddingRight(dp int) *FTextView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FTextView) PaddingBottom(dp int) *FTextView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FTextView) PaddingAll(all int) *FTextView {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FTextView) Margin(left, top, right, bottom int) *FTextView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FTextView) MarginLeft(dp int) *FTextView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FTextView) MarginTop(dp int) *FTextView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FTextView) MarginRight(dp int) *FTextView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FTextView) MarginBottom(dp int) *FTextView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FTextView) MarginAll(dp int) *FTextView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FTextView) LayoutGravity(gravity int) *FTextView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FTextView) Elevation(dp float64) *FTextView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FTextView) Assign(fb **FTextView) *FTextView {
	*fb = v
	return v
}
func (v *FTextView) LayoutWeight(f int) *FTextView {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FTextView) OnTouch(f func(TouchEvent)) *FTextView {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FTextView) OnClick(f func()) *FTextView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

// --------------------------------------------------------
func (v *FTextView) Text(s string) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}

func (v *FTextView) TextColor(s string) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}

func (v *FTextView) TextSize(dpsize int) *FTextView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextSize", SPrintf(dpsize))
	return v
}
func (v *FTextView) GetText() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Text")
}
