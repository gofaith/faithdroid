package faithdroid

type FListView struct {
	FBaseView
	Vh ViewHolder
}
type TypeOnBindDataArgsBundle struct {
	Str      string
	Position int
}

// ViewHolder
type ViewHolder struct {
	VID   string
	Vlist map[string]string
}

func (vh *ViewHolder) GetListViewByItemId(id string) *FListView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FListView); ok {
			return bt
		}
	}
	return nil
}

// -----------------------------------
func (v *FListView) SetId(s string) *FListView {
	GlobalVars.IdMap[s] = v
	return v
}
func (v *FListView) SetItemId(parent *FListView, id string) *FListView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetListViewById(id string) *FListView {
	if v, ok := GlobalVars.IdMap[id].(*FListView); ok {
		return v
	}
	return nil
}
func VListView(a IActivity, getView func(lv *FListView) IView, bindData func(vh *ViewHolder, pos int), getCount func() int) *FListView {
	v := &FListView{}
	v.VID = NewToken()
	v.ClassName = "VListView"
	v.UI = a.GetMyActivity().UI
	v.Vh.Vlist = make(map[string]string)
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	fnId1 := NewToken()
	GlobalVars.EventHandlersMap[fnId1] = func(string) string {
		v.Vh.VID = getView(v).GetViewId()
		return JsonObject(v.Vh)
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetView", fnId1)

	fnId2 := NewToken()
	GlobalVars.EventHandlersMap[fnId2] = func(str string) string {
		obd := TypeOnBindDataArgsBundle{}
		UnJson(str, &obd)
		vh := ViewHolder{}
		UnJson(obd.Str, &vh)
		bindData(&vh, obd.Position)
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnBindData", fnId2)

	fnId3 := NewToken()
	GlobalVars.EventHandlersMap[fnId3] = func(str string) string {
		return SPrintf(getCount())
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId3)
	return v
}
func HListView(a IActivity, getView func(lv *FListView) IView, bindData func(vh *ViewHolder, pos int), getCount func() int) *FListView {
	v := &FListView{}
	v.VID = NewToken()
	v.ClassName = "HListView"
	v.UI = a.GetMyActivity().UI
	v.Vh.Vlist = make(map[string]string)
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	fnId1 := NewToken()
	GlobalVars.EventHandlersMap[fnId1] = func(string) string {
		v.Vh.VID = getView(v).GetViewId()
		return JsonObject(v.Vh)
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetView", fnId1)

	fnId2 := NewToken()
	GlobalVars.EventHandlersMap[fnId2] = func(str string) string {
		obd := TypeOnBindDataArgsBundle{}
		UnJson(str, &obd)
		vh := ViewHolder{}
		UnJson(obd.Str, &vh)
		bindData(&vh, obd.Position)
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnBindData", fnId2)

	fnId3 := NewToken()
	GlobalVars.EventHandlersMap[fnId3] = func(str string) string {
		return SPrintf(getCount())
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId3)
	return v
}
func (v *FListView) Size(w, h int) *FListView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FListView) X(x float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FListView) Y(y float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FListView) PivotX(x float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FListView) PivotY(y float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FListView) ScaleX(x float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FListView) ScaleY(y float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FListView) Rotation(r float64) *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FListView) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FListView) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FListView) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FListView) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FListView) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FListView) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FListView) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FListView) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FListView) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FListView) Background(s string) *FListView {
	v.FBaseView.Background(s)
	return v
}
func (v *FListView) Foreground(s string) *FListView {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FListView) CachedForeground(s string) *FListView {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FListView) BackgroundColor(s string) *FListView {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FListView) CachedBackground(s string) *FListView {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FListView) onClick(f func()) *FListView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FListView) Visible() *FListView {
	v.FBaseView.Visible()
	return v
}
func (v *FListView) Invisible() *FListView {
	v.FBaseView.Invisible()
	return v
}
func (v *FListView) Gone() *FListView {
	v.FBaseView.Gone()
	return v
}

func (v *FListView) Padding(left, top, right, bottom int) *FListView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FListView) PaddingLeft(dp int) *FListView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FListView) PaddingTop(dp int) *FListView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FListView) PaddingRight(dp int) *FListView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FListView) PaddingBottom(dp int) *FListView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FListView) PaddingAll(all int) *FListView {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FListView) Margin(left, top, right, bottom int) *FListView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FListView) MarginLeft(dp int) *FListView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FListView) MarginTop(dp int) *FListView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FListView) MarginRight(dp int) *FListView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FListView) MarginBottom(dp int) *FListView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FListView) MarginAll(dp int) *FListView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FListView) LayoutGravity(gravity int) *FListView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FListView) Elevation(dp float64) *FListView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FListView) Assign(fb **FListView) *FListView {
	*fb = v
	return v
}
func (v *FListView) LayoutWeight(f int) *FListView {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FListView) OnTouch(f func(TouchEvent)) *FListView {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FListView) OnClick(f func()) *FListView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FListView) Clickable(b bool) *FListView {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FListView) Top2TopOf(id string) *FListView {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FListView) Top2BottomOf(id string) *FListView {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FListView) Bottom2TopOf(id string) *FListView {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FListView) Bottom2BottomOf(id string) *FListView {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FListView) Left2LeftOf(id string) *FListView {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FListView) Right2RightOf(id string) *FListView {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FListView) Left2RightOf(id string) *FListView {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FListView) Right2LeftOf(id string) *FListView {
	v.FBaseView.Right2LeftOf(id)
	return v
}

// --------------------------------------------------------------------
func (v *FListView) NotifyDataSetChanged() *FListView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "NotifyDataSetChanged", "")
	return v
}
