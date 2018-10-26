package faithdroid

type FCheckBox struct {
	FBaseView
}

func CheckBox(a IActivity) *FCheckBox {
	v := &FCheckBox{}
	v.VID = NewToken()
	v.ClassName = "CheckBox"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetCheckBoxByItemId(id string) *FCheckBox {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FCheckBox); ok {
			return bt
		}
	}
	return nil
}

func (v *FCheckBox) Size(w, h int) *FCheckBox {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FCheckBox) X(x float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FCheckBox) Y(y float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FCheckBox) PivotX(x float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FCheckBox) PivotY(y float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FCheckBox) ScaleX(x float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FCheckBox) ScaleY(y float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FCheckBox) Rotation(r float64) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FCheckBox) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FCheckBox) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FCheckBox) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FCheckBox) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FCheckBox) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FCheckBox) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FCheckBox) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FCheckBox) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FCheckBox) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FCheckBox) SetId(s string) *FCheckBox {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FCheckBox) SetItemId(parent *FListView, id string) *FCheckBox {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetCheckBoxById(id string) *FCheckBox {
	if v, ok := GlobalVars.IdMap[id].(*FCheckBox); ok {
		return v
	}
	return nil
}

func (v *FCheckBox) Background(s string) *FCheckBox {
	v.FBaseView.Background(s)
	return v
}
func (v *FCheckBox) Foreground(s string) *FCheckBox {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FCheckBox) CachedForeground(s string) *FCheckBox {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FCheckBox) BackgroundColor(s string) *FCheckBox {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FCheckBox) CachedBackground(s string) *FCheckBox {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FCheckBox) Visible() *FCheckBox {
	v.FBaseView.Visible()
	return v
}
func (v *FCheckBox) Invisible() *FCheckBox {
	v.FBaseView.Invisible()
	return v
}
func (v *FCheckBox) Gone() *FCheckBox {
	v.FBaseView.Gone()
	return v
}

func (v *FCheckBox) Padding(left, top, right, bottom int) *FCheckBox {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FCheckBox) PaddingLeft(dp int) *FCheckBox {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FCheckBox) PaddingTop(dp int) *FCheckBox {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FCheckBox) PaddingRight(dp int) *FCheckBox {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FCheckBox) PaddingBottom(dp int) *FCheckBox {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FCheckBox) PaddingAll(dp int) *FCheckBox {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FCheckBox) Margin(left, top, right, bottom int) *FCheckBox {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FCheckBox) MarginLeft(dp int) *FCheckBox {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FCheckBox) MarginTop(dp int) *FCheckBox {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FCheckBox) MarginRight(dp int) *FCheckBox {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FCheckBox) MarginBottom(dp int) *FCheckBox {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FCheckBox) MarginAll(dp int) *FCheckBox {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FCheckBox) LayoutGravity(gravity int) *FCheckBox {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FCheckBox) Elevation(dp float64) *FCheckBox {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FCheckBox) Assign(fb **FCheckBox) *FCheckBox {
	*fb = v
	return v
}
func (v *FCheckBox) LayoutWeight(f int) *FCheckBox {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FCheckBox) OnTouch(f func(TouchEvent)) *FCheckBox {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FCheckBox) OnClick(f func()) *FCheckBox {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FCheckBox) Clickable(b bool) *FCheckBox {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FCheckBox) Top2TopOf(id string) *FCheckBox {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FCheckBox) Top2BottomOf(id string) *FCheckBox {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FCheckBox) Bottom2TopOf(id string) *FCheckBox {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FCheckBox) Bottom2BottomOf(id string) *FCheckBox {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FCheckBox) Left2LeftOf(id string) *FCheckBox {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FCheckBox) Right2RightOf(id string) *FCheckBox {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FCheckBox) Left2RightOf(id string) *FCheckBox {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FCheckBox) Right2LeftOf(id string) *FCheckBox {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FCheckBox) CenterX() *FCheckBox {
	v.FBaseView.CenterX()
	return v
}
func (v *FCheckBox) CenterY() *FCheckBox {
	v.FBaseView.CenterY()
	return v
}
func (v *FCheckBox) WidthPercent(num float64) *FCheckBox {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FCheckBox) HeightPercent(num float64) *FCheckBox {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
func (v *FCheckBox) Text(s string) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FCheckBox) GetText() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FCheckBox) Enabled(b bool) *FCheckBox {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FCheckBox) IsEnabled() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}
func (v *FCheckBox) Checked(b bool) *FCheckBox {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Checked", SPrintf(b))
	return v
}
func (v *FCheckBox) IsChecked() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Checked") == "true" {
		return true
	}
	return false
}
func (v *FCheckBox) OnCheckedChangeListener(f func(bool)) *FCheckBox {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(s string) string {
		if s == "true" {
			f(true)
		} else {
			f(false)
		}
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnCheckedChangeListener", fnID)
	return v
}
