package faithdroid

type FRadioButton struct {
	FBaseView
}

func RadioButton(a *Activity) *FRadioButton {
	v := &FRadioButton{}
	v.VID = NewToken()
	v.ClassName = "RadioButton"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.BaseMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetRadioButtonByItemId(id string) *FRadioButton {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.BaseMap[v].(*FRadioButton); ok {
			return bt
		}
	}
	return nil
}

func (v *FRadioButton) Size(w, h int) *FRadioButton {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FRadioButton) X(x float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FRadioButton) Y(y float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FRadioButton) PivotX(x float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FRadioButton) PivotY(y float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FRadioButton) ScaleX(x float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FRadioButton) ScaleY(y float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FRadioButton) Rotation(r float64) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FRadioButton) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FRadioButton) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FRadioButton) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FRadioButton) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FRadioButton) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FRadioButton) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FRadioButton) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FRadioButton) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FRadioButton) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FRadioButton) SetId(s string) *FRadioButton {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FRadioButton) SetItemId(parent *FListView, id string) *FRadioButton {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetRadioButtonById(id string) *FRadioButton {
	if v, ok := GlobalVars.IdMap[id].(*FRadioButton); ok {
		return v
	}
	return nil
}

func (v *FRadioButton) Background(s string) *FRadioButton {
	v.FBaseView.Background(s)
	return v
}
func (v *FRadioButton) Foreground(s string) *FRadioButton {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FRadioButton) CachedForeground(s string) *FRadioButton {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FRadioButton) BackgroundColor(s string) *FRadioButton {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FRadioButton) CachedBackground(s string) *FRadioButton {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FRadioButton) Visible() *FRadioButton {
	v.FBaseView.Visible()
	return v
}
func (v *FRadioButton) Invisible() *FRadioButton {
	v.FBaseView.Invisible()
	return v
}
func (v *FRadioButton) Gone() *FRadioButton {
	v.FBaseView.Gone()
	return v
}

func (v *FRadioButton) Padding(left, top, right, bottom int) *FRadioButton {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FRadioButton) PaddingLeft(dp int) *FRadioButton {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FRadioButton) PaddingTop(dp int) *FRadioButton {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FRadioButton) PaddingRight(dp int) *FRadioButton {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FRadioButton) PaddingBottom(dp int) *FRadioButton {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FRadioButton) PaddingAll(dp int) *FRadioButton {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FRadioButton) Margin(left, top, right, bottom int) *FRadioButton {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FRadioButton) MarginLeft(dp int) *FRadioButton {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FRadioButton) MarginTop(dp int) *FRadioButton {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FRadioButton) MarginRight(dp int) *FRadioButton {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FRadioButton) MarginBottom(dp int) *FRadioButton {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FRadioButton) MarginAll(dp int) *FRadioButton {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FRadioButton) LayoutGravity(gravity int) *FRadioButton {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FRadioButton) Elevation(dp float64) *FRadioButton {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FRadioButton) Assign(fb **FRadioButton) *FRadioButton {
	*fb = v
	return v
}
func (v *FRadioButton) LayoutWeight(f int) *FRadioButton {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FRadioButton) OnTouch(f func(TouchEvent)) *FRadioButton {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FRadioButton) OnClick(f func()) *FRadioButton {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FRadioButton) TopToTop(iv IBaseView) *FRadioButton {
	v.FBaseView.TopToTop(iv)
	return v
}
func (v *FRadioButton) TopToBottom(iv IBaseView) *FRadioButton {
	v.FBaseView.TopToBottom(iv)
	return v
}
func (v *FRadioButton) BottomToBottom(iv IBaseView) *FRadioButton {
	v.FBaseView.BottomToBottom(iv)
	return v
}
func (v *FRadioButton) BottomToTop(iv IBaseView) *FRadioButton {
	v.FBaseView.BottomToTop(iv)
	return v
}
func (v *FRadioButton) LeftToLeft(iv IBaseView) *FRadioButton {
	v.FBaseView.LeftToLeft(iv)
	return v
}
func (v *FRadioButton) LeftToRight(iv IBaseView) *FRadioButton {
	v.FBaseView.LeftToRight(iv)
	return v
}
func (v *FRadioButton) RightToRight(iv IBaseView) *FRadioButton {
	v.FBaseView.RightToRight(iv)
	return v
}
func (v *FRadioButton) RightToLeft(iv IBaseView) *FRadioButton {
	v.FBaseView.RightToLeft(iv)
	return v
}
func (v *FRadioButton) GetBaseView() *FBaseView {
	return &v.FBaseView
}
// --------------------------------------------------------
func (v *FRadioButton) Text(s string) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FRadioButton) TextColor(s string) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}
func (v *FRadioButton) TextSize(dpsize int) *FRadioButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextSize", SPrintf(dpsize))
	return v
}
func (v *FRadioButton) GetText() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FRadioButton) Enabled(b bool) *FRadioButton {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FRadioButton) IsEnabled() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}
