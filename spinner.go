package faithdroid

type FSpinner struct {
	FBaseView
}

func Spinner(a *Activity) *FSpinner {
	v := &FSpinner{}
	v.VID = NewToken()
	v.ClassName = "Spinner"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetSpinnerByItemId(id string) *FSpinner {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FSpinner); ok {
			return bt
		}
	}
	return nil
}

func (v *FSpinner) Size(w, h int) *FSpinner {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FSpinner) X(x float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FSpinner) Y(y float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FSpinner) PivotX(x float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FSpinner) PivotY(y float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FSpinner) ScaleX(x float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FSpinner) ScaleY(y float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FSpinner) Rotation(r float64) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FSpinner) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FSpinner) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FSpinner) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FSpinner) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FSpinner) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FSpinner) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FSpinner) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FSpinner) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FSpinner) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FSpinner) SetId(s string) *FSpinner {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FSpinner) SetItemId(parent *FListView, id string) *FSpinner {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetSpinnerById(id string) *FSpinner {
	if v, ok := GlobalVars.IdMap[id].(*FSpinner); ok {
		return v
	}
	return nil
}

func (v *FSpinner) Background(s string) *FSpinner {
	v.FBaseView.Background(s)
	return v
}
func (v *FSpinner) Foreground(s string) *FSpinner {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FSpinner) CachedForeground(s string) *FSpinner {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FSpinner) BackgroundColor(s string) *FSpinner {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FSpinner) CachedBackground(s string) *FSpinner {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FSpinner) Visible() *FSpinner {
	v.FBaseView.Visible()
	return v
}
func (v *FSpinner) Invisible() *FSpinner {
	v.FBaseView.Invisible()
	return v
}
func (v *FSpinner) Gone() *FSpinner {
	v.FBaseView.Gone()
	return v
}

func (v *FSpinner) Padding(left, top, right, bottom int) *FSpinner {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FSpinner) PaddingLeft(dp int) *FSpinner {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FSpinner) PaddingTop(dp int) *FSpinner {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FSpinner) PaddingRight(dp int) *FSpinner {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FSpinner) PaddingBottom(dp int) *FSpinner {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FSpinner) PaddingAll(dp int) *FSpinner {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FSpinner) Margin(left, top, right, bottom int) *FSpinner {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FSpinner) MarginLeft(dp int) *FSpinner {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FSpinner) MarginTop(dp int) *FSpinner {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FSpinner) MarginRight(dp int) *FSpinner {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FSpinner) MarginBottom(dp int) *FSpinner {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FSpinner) MarginAll(dp int) *FSpinner {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FSpinner) LayoutGravity(gravity int) *FSpinner {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FSpinner) Elevation(dp float64) *FSpinner {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FSpinner) Assign(fb **FSpinner) *FSpinner {
	*fb = v
	return v
}
func (v *FSpinner) LayoutWeight(f int) *FSpinner {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FSpinner) OnTouch(f func(TouchEvent)) *FSpinner {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FSpinner) OnClick(f func()) *FSpinner {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

// --------------------------------------------------------

func (v *FSpinner) Enabled(b bool) *FSpinner {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Enabled", "false")
	}
	return v
}
func (v *FSpinner) IsEnabled() bool {
	if GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Enabled") == "true" {
		return true
	}
	return false
}
func (v *FSpinner) NotifyDataSetChanged() *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "NotifyDataSetChanged", "")
	return v
}
func (v *FSpinner) List(ls []string) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "List", JsonArray(ls))
	return v
}

func (v *FSpinner) ListAdd(s string) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ListAdd", s)
	return v
}

func (v *FSpinner) ListRemove(i int) *FSpinner {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ListRemove", SPrintf(i))
	return v
}

func (v *FSpinner) OnItemClick(f func(int)) *FSpinner {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(s string) string {
		i, e := a2i(s)
		if e != nil {
			return ""
		}
		f(i)
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnItemClick", fnId)
	return v
}
