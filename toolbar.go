package faithdroid

type FToolbar struct {
	FBaseView
	MyMenu []interface{}
}

// --------------------------------------------------------------------------------
func Toolbar(a *Activity) *FToolbar {
	v := &FToolbar{}
	v.VID = NewToken()
	v.ClassName = "Toolbar"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FToolbar) Size(w, h int) *FToolbar {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FToolbar) X(x float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FToolbar) Y(y float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FToolbar) PivotX(x float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FToolbar) PivotY(y float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FToolbar) ScaleX(x float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FToolbar) ScaleY(y float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FToolbar) Rotation(r float64) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FToolbar) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FToolbar) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FToolbar) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FToolbar) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FToolbar) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FToolbar) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FToolbar) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FToolbar) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FToolbar) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FToolbar) SetId(s string) *FToolbar {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FToolbar) SetItemId(parent *FListView, id string) *FToolbar {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetToolbarById(id string) *FToolbar {
	if v, ok := GlobalVars.IdMap[id].(*FToolbar); ok {
		return v
	}
	return nil
}

func (v *FToolbar) Background(s string) *FToolbar {
	v.FBaseView.Background(s)
	return v
}
func (v *FToolbar) Foreground(s string) *FToolbar {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FToolbar) CachedForeground(s string) *FToolbar {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FToolbar) BackgroundColor(s string) *FToolbar {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FToolbar) CachedBackground(s string) *FToolbar {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FToolbar) Visible() *FToolbar {
	v.FBaseView.Visible()
	return v
}
func (v *FToolbar) Invisible() *FToolbar {
	v.FBaseView.Invisible()
	return v
}
func (v *FToolbar) Gone() *FToolbar {
	v.FBaseView.Gone()
	return v
}

func (v *FToolbar) Padding(left, top, right, bottom int) *FToolbar {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FToolbar) PaddingLeft(dp int) *FToolbar {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) PaddingTop(dp int) *FToolbar {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FToolbar) PaddingRight(dp int) *FToolbar {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FToolbar) PaddingBottom(dp int) *FToolbar {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FToolbar) PaddingAll(dp int) *FToolbar {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) Margin(left, top, right, bottom int) *FToolbar {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FToolbar) MarginLeft(dp int) *FToolbar {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) MarginTop(dp int) *FToolbar {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FToolbar) MarginRight(dp int) *FToolbar {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FToolbar) MarginBottom(dp int) *FToolbar {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FToolbar) MarginAll(dp int) *FToolbar {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) LayoutGravity(gravity int) *FToolbar {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FToolbar) Elevation(dp float64) *FToolbar {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FToolbar) Assign(fb **FToolbar) *FToolbar {
	*fb = v
	return v
}
func (v *FToolbar) LayoutWeight(f int) *FToolbar {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FToolbar) OnTouch(f func(TouchEvent)) *FToolbar {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FToolbar) OnClick(f func()) *FToolbar {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

// ----------------------------------------------------------
func (v *FToolbar) Title(s string) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FToolbar) Subtitle(s string) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "SubTitle", s)
	return v
}
func (v *FToolbar) Menus(mis ...interface{}) *FToolbar {
	v.MyMenu = mis
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Menu", JsonArray(v.MyMenu))
	return v
}
