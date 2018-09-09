package faithdroid

type FButton struct {
	FBaseView
}

func Button(a *Activity) *FButton {
	v := &FButton{}
	v.VID = NewToken()
	v.ClassName = "Button"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetButtonByItemId(id string) *FButton {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FButton); ok {
			return bt
		}
	}
	return nil
}

func (v *FButton) Size(w, h int) *FButton {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FButton) X(x float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FButton) Y(y float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FButton) PivotX(x float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FButton) PivotY(y float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FButton) ScaleX(x float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FButton) ScaleY(y float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FButton) Rotation(r float64) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FButton) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FButton) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FButton) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FButton) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FButton) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FButton) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FButton) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FButton) SetId(s string) *FButton {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FButton) SetItemId(parent *FListView, id string) *FButton {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetButtonById(id string) *FButton {
	if v, ok := GlobalVars.IdMap[id].(*FButton); ok {
		return v
	}
	return nil
}

func (v *FButton) Background(s string) *FButton {
	v.FBaseView.Background(s)
	return v
}
func (v *FButton) BackgroundColor(s string) *FButton {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FButton) CachedBackground(s string) *FButton {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FButton) Visible() *FButton {
	v.FBaseView.Visible()
	return v
}
func (v *FButton) Invisible() *FButton {
	v.FBaseView.Invisible()
	return v
}
func (v *FButton) Gone() *FButton {
	v.FBaseView.Gone()
	return v
}

func (v *FButton) Padding(left, top, right, bottom int) *FButton {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FButton) PaddingLeft(dp int) *FButton {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FButton) PaddingTop(dp int) *FButton {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FButton) PaddingRight(dp int) *FButton {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FButton) PaddingBottom(dp int) *FButton {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FButton) PaddingAll(dp int) *FButton {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FButton) Margin(left, top, right, bottom int) *FButton {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FButton) MarginLeft(dp int) *FButton {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FButton) MarginTop(dp int) *FButton {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FButton) MarginRight(dp int) *FButton {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FButton) MarginBottom(dp int) *FButton {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FButton) MarginAll(dp int) *FButton {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FButton) LayoutGravity(gravity int) *FButton {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FButton) Elevation(dp float64) *FButton {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FButton) Assign(fb **FButton) *FButton {
	*fb = v
	return v
}
func (v *FButton) LayoutWeight(f int) *FButton {
	v.FBaseView.LayoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FButton) Text(s string) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FButton) TextColor(s string) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}
func (v *FButton) TextSize(dpsize int) *FButton {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextSize", SPrintf(dpsize))
	return v
}
func (v *FButton) GetText() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FButton) Enabled(b bool) *FButton {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FButton) IsEnabled() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}

func (v *FButton) OnClick(f func()) *FButton {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
