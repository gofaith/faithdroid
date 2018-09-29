package faithdroid

type FFab struct {
	FBaseView
}

func Fab(a IActivity) *FFab {
	v := &FFab{}
	v.VID = NewToken()
	v.ClassName = "Fab"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetFabByItemId(id string) *FFab {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FFab); ok {
			return bt
		}
	}
	return nil
}

func (v *FFab) Size(w, h int) *FFab {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FFab) X(x float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FFab) Y(y float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FFab) PivotX(x float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FFab) PivotY(y float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FFab) ScaleX(x float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FFab) ScaleY(y float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FFab) Rotation(r float64) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FFab) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FFab) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FFab) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FFab) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FFab) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FFab) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FFab) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FFab) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FFab) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FFab) SetId(s string) *FFab {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FFab) SetItemId(parent *FListView, id string) *FFab {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetFabById(id string) *FFab {
	if v, ok := GlobalVars.IdMap[id].(*FFab); ok {
		return v
	}
	return nil
}

func (v *FFab) Background(s string) *FFab {
	v.FBaseView.Background(s)
	return v
}
func (v *FFab) Foreground(s string) *FFab {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FFab) CachedForeground(s string) *FFab {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FFab) BackgroundColor(s string) *FFab {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FFab) CachedBackground(s string) *FFab {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FFab) Visible() *FFab {
	v.FBaseView.Visible()
	return v
}
func (v *FFab) Invisible() *FFab {
	v.FBaseView.Invisible()
	return v
}
func (v *FFab) Gone() *FFab {
	v.FBaseView.Gone()
	return v
}

func (v *FFab) Padding(left, top, right, bottom int) *FFab {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FFab) PaddingLeft(dp int) *FFab {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FFab) PaddingTop(dp int) *FFab {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FFab) PaddingRight(dp int) *FFab {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FFab) PaddingBottom(dp int) *FFab {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FFab) PaddingAll(dp int) *FFab {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FFab) Margin(left, top, right, bottom int) *FFab {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FFab) MarginLeft(dp int) *FFab {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FFab) MarginTop(dp int) *FFab {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FFab) MarginRight(dp int) *FFab {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FFab) MarginBottom(dp int) *FFab {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FFab) MarginAll(dp int) *FFab {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FFab) LayoutGravity(gravity int) *FFab {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FFab) Elevation(dp float64) *FFab {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FFab) Assign(fb **FFab) *FFab {
	*fb = v
	return v
}
func (v *FFab) LayoutWeight(f int) *FFab {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FFab) OnTouch(f func(TouchEvent)) *FFab {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FFab) Clickable(b bool) *FFab {
	v.FBaseView.Clickable(b)
	return v
}

// --------------------------------------------------------

func (v *FFab) Icon(s string) *FFab {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Icon", s)
	return v
}
func (v *FFab) OnClick(f func()) *FFab {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnId)
	return v
}
