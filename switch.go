package faithdroid

type FSwitch struct {
	FBaseView
}

func Switch(a IActivity) *FSwitch {
	v := &FSwitch{}
	v.VID = NewToken()
	v.ClassName = "Switch"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetSwitchByItemId(id string) *FSwitch {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FSwitch); ok {
			return bt
		}
	}
	return nil
}

func (v *FSwitch) Size(w, h int) *FSwitch {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FSwitch) X(x float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FSwitch) Y(y float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FSwitch) PivotX(x float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FSwitch) PivotY(y float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FSwitch) ScaleX(x float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FSwitch) ScaleY(y float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FSwitch) Rotation(r float64) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FSwitch) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FSwitch) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FSwitch) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FSwitch) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FSwitch) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FSwitch) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FSwitch) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FSwitch) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FSwitch) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FSwitch) SetId(s string) *FSwitch {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FSwitch) SetItemId(parent *FListView, id string) *FSwitch {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetSwitchById(id string) *FSwitch {
	if v, ok := GlobalVars.IdMap[id].(*FSwitch); ok {
		return v
	}
	return nil
}

func (v *FSwitch) Background(s string) *FSwitch {
	v.FBaseView.Background(s)
	return v
}
func (v *FSwitch) Foreground(s string) *FSwitch {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FSwitch) CachedForeground(s string) *FSwitch {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FSwitch) BackgroundColor(s string) *FSwitch {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FSwitch) CachedBackground(s string) *FSwitch {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FSwitch) Visible() *FSwitch {
	v.FBaseView.Visible()
	return v
}
func (v *FSwitch) Invisible() *FSwitch {
	v.FBaseView.Invisible()
	return v
}
func (v *FSwitch) Gone() *FSwitch {
	v.FBaseView.Gone()
	return v
}

func (v *FSwitch) Padding(left, top, right, bottom int) *FSwitch {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FSwitch) PaddingLeft(dp int) *FSwitch {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FSwitch) PaddingTop(dp int) *FSwitch {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FSwitch) PaddingRight(dp int) *FSwitch {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FSwitch) PaddingBottom(dp int) *FSwitch {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FSwitch) PaddingAll(dp int) *FSwitch {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FSwitch) Margin(left, top, right, bottom int) *FSwitch {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FSwitch) MarginLeft(dp int) *FSwitch {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FSwitch) MarginTop(dp int) *FSwitch {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FSwitch) MarginRight(dp int) *FSwitch {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FSwitch) MarginBottom(dp int) *FSwitch {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FSwitch) MarginAll(dp int) *FSwitch {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FSwitch) LayoutGravity(gravity int) *FSwitch {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FSwitch) Elevation(dp float64) *FSwitch {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FSwitch) Assign(fb **FSwitch) *FSwitch {
	*fb = v
	return v
}
func (v *FSwitch) LayoutWeight(f int) *FSwitch {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FSwitch) OnTouch(f func(TouchEvent)) *FSwitch {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FSwitch) OnClick(f func()) *FSwitch {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FSwitch) Clickable(b bool) *FSwitch {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FSwitch) Top2TopOf(id string) *FSwitch {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FSwitch) Top2BottomOf(id string) *FSwitch {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FSwitch) Bottom2TopOf(id string) *FSwitch {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FSwitch) Bottom2BottomOf(id string) *FSwitch {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FSwitch) Left2LeftOf(id string) *FSwitch {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FSwitch) Right2RightOf(id string) *FSwitch {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FSwitch) Left2RightOf(id string) *FSwitch {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FSwitch) Right2LeftOf(id string) *FSwitch {
	v.FBaseView.Right2LeftOf(id)
	return v
}

// --------------------------------------------------------
func (v *FSwitch) Enabled(b bool) *FSwitch {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FSwitch) IsEnabled() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}
func (v *FSwitch) Checked(b bool) *FSwitch {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Checked", SPrintf(b))
	return v
}
func (v *FSwitch) IsChecked() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Checked") == "true" {
		return true
	}
	return false
}
func (v *FSwitch) OnCheckedChangeListener(f func(bool)) *FSwitch {
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
