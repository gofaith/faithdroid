package faithdroid

type FBottomNav struct {
	FBaseView
	MyMenu []interface{}
}

func (vh *ViewHolder) GetFBottomNavByItemId(id string) *FBottomNav {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FBottomNav); ok {
			return bt
		}
	}
	return nil
}
func (v *FBottomNav) SetId(s string) *FBottomNav {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FBottomNav) SetItemId(parent *FListView, id string) *FBottomNav {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetBottomNavById(id string) *FBottomNav {
	if v, ok := GlobalVars.IdMap[id].(*FBottomNav); ok {
		return v
	}
	return nil
}
func BottomNav(a IActivity) *FBottomNav {
	v := &FBottomNav{}
	v.VID = NewToken()
	v.ClassName = "BottomNav"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FBottomNav) Size(w, h int) *FBottomNav {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FBottomNav) X(x float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FBottomNav) Y(y float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FBottomNav) PivotX(x float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FBottomNav) PivotY(y float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FBottomNav) ScaleX(x float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FBottomNav) ScaleY(y float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FBottomNav) Rotation(r float64) *FBottomNav {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FBottomNav) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FBottomNav) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FBottomNav) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FBottomNav) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FBottomNav) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FBottomNav) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FBottomNav) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FBottomNav) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FBottomNav) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FBottomNav) Background(s string) *FBottomNav {
	v.FBaseView.Background(s)
	return v
}
func (v *FBottomNav) Foreground(s string) *FBottomNav {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FBottomNav) CachedForeground(s string) *FBottomNav {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FBottomNav) BackgroundColor(s string) *FBottomNav {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FBottomNav) CachedBackground(s string) *FBottomNav {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FBottomNav) onClick(f func()) *FBottomNav {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FBottomNav) Visible() *FBottomNav {
	v.FBaseView.Visible()
	return v
}
func (v *FBottomNav) Invisible() *FBottomNav {
	v.FBaseView.Invisible()
	return v
}
func (v *FBottomNav) Gone() *FBottomNav {
	v.FBaseView.Gone()
	return v
}

func (v *FBottomNav) Padding(left, top, right, bottom int) *FBottomNav {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FBottomNav) PaddingLeft(dp int) *FBottomNav {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FBottomNav) PaddingTop(dp int) *FBottomNav {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FBottomNav) PaddingRight(dp int) *FBottomNav {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FBottomNav) PaddingBottom(dp int) *FBottomNav {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FBottomNav) PaddingAll(all int) *FBottomNav {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FBottomNav) Margin(left, top, right, bottom int) *FBottomNav {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FBottomNav) MarginLeft(dp int) *FBottomNav {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FBottomNav) MarginTop(dp int) *FBottomNav {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FBottomNav) MarginRight(dp int) *FBottomNav {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FBottomNav) MarginBottom(dp int) *FBottomNav {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FBottomNav) MarginAll(dp int) *FBottomNav {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FBottomNav) LayoutGravity(gravity int) *FBottomNav {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FBottomNav) Elevation(dp float64) *FBottomNav {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FBottomNav) Assign(fb **FBottomNav) *FBottomNav {
	*fb = v
	return v
}
func (v *FBottomNav) LayoutWeight(f int) *FBottomNav {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FBottomNav) OnTouch(f func(TouchEvent)) *FBottomNav {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FBottomNav) OnClick(f func()) *FBottomNav {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

// --------------------------------------------------------
func (v *FBottomNav) Menus(mis ...interface{}) *FBottomNav {
	v.MyMenu = mis
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Menu", JsonArray(v.MyMenu))
	return v
}
