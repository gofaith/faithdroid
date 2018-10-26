package faithdroid

type FRadioGroup struct {
	FBaseView
	Children []*FRadioButton
}

func (vh *ViewHolder) GetRadioGroupByItemId(id string) *FRadioGroup {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FRadioGroup); ok {
			return bt
		}
	}
	return nil
}
func (v *FRadioGroup) SetId(s string) *FRadioGroup {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FRadioGroup) SetItemId(parent *FListView, id string) *FRadioGroup {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetRadioGroupById(id string) *FRadioGroup {
	if v, ok := GlobalVars.IdMap[id].(*FRadioGroup); ok {
		return v
	}
	return nil
}
func RadioGroup(a IActivity) *FRadioGroup {
	v := &FRadioGroup{}
	v.VID = NewToken()
	v.ClassName = "RadioGroup"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FRadioGroup) Size(w, h int) *FRadioGroup {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FRadioGroup) X(x float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FRadioGroup) Y(y float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FRadioGroup) PivotX(x float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FRadioGroup) PivotY(y float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FRadioGroup) ScaleX(x float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FRadioGroup) ScaleY(y float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FRadioGroup) Rotation(r float64) *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FRadioGroup) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FRadioGroup) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FRadioGroup) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FRadioGroup) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FRadioGroup) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FRadioGroup) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FRadioGroup) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FRadioGroup) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FRadioGroup) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}

func (v *FRadioGroup) Background(s string) *FRadioGroup {
	v.FBaseView.Background(s)
	return v
}
func (v *FRadioGroup) Foreground(s string) *FRadioGroup {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FRadioGroup) CachedForeground(s string) *FRadioGroup {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FRadioGroup) BackgroundColor(s string) *FRadioGroup {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FRadioGroup) CachedBackground(s string) *FRadioGroup {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FRadioGroup) onClick(f func()) *FRadioGroup {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FRadioGroup) Visible() *FRadioGroup {
	v.FBaseView.Visible()
	return v
}
func (v *FRadioGroup) Invisible() *FRadioGroup {
	v.FBaseView.Invisible()
	return v
}
func (v *FRadioGroup) Gone() *FRadioGroup {
	v.FBaseView.Gone()
	return v
}

func (v *FRadioGroup) Padding(left, top, right, bottom int) *FRadioGroup {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FRadioGroup) PaddingLeft(dp int) *FRadioGroup {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FRadioGroup) PaddingTop(dp int) *FRadioGroup {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FRadioGroup) PaddingRight(dp int) *FRadioGroup {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FRadioGroup) PaddingBottom(dp int) *FRadioGroup {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FRadioGroup) PaddingAll(all int) *FRadioGroup {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FRadioGroup) Margin(left, top, right, bottom int) *FRadioGroup {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FRadioGroup) MarginLeft(dp int) *FRadioGroup {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FRadioGroup) MarginTop(dp int) *FRadioGroup {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FRadioGroup) MarginRight(dp int) *FRadioGroup {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FRadioGroup) MarginBottom(dp int) *FRadioGroup {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FRadioGroup) MarginAll(dp int) *FRadioGroup {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FRadioGroup) LayoutGravity(gravity int) *FRadioGroup {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FRadioGroup) Elevation(dp float64) *FRadioGroup {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FRadioGroup) Assign(fb **FRadioGroup) *FRadioGroup {
	*fb = v
	return v
}
func (v *FRadioGroup) LayoutWeight(f int) *FRadioGroup {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FRadioGroup) OnTouch(f func(TouchEvent)) *FRadioGroup {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FRadioGroup) OnClick(f func()) *FRadioGroup {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FRadioGroup) Clickable(b bool) *FRadioGroup {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FRadioGroup) Top2TopOf(id string) *FRadioGroup {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FRadioGroup) Top2BottomOf(id string) *FRadioGroup {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FRadioGroup) Bottom2TopOf(id string) *FRadioGroup {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FRadioGroup) Bottom2BottomOf(id string) *FRadioGroup {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FRadioGroup) Left2LeftOf(id string) *FRadioGroup {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FRadioGroup) Right2RightOf(id string) *FRadioGroup {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FRadioGroup) Left2RightOf(id string) *FRadioGroup {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FRadioGroup) Right2LeftOf(id string) *FRadioGroup {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FRadioGroup) CenterX() *FRadioGroup {
	v.FBaseView.CenterX()
	return v
}
func (v *FRadioGroup) CenterY() *FRadioGroup {
	v.FBaseView.CenterY()
	return v
}
func (v *FRadioGroup) WidthPercent(num float64) *FRadioGroup {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FRadioGroup) HeightPercent(num float64) *FRadioGroup {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
func (v *FRadioGroup) Append(vs ...*FRadioButton) *FRadioGroup {
	v.Children = append(v.Children, vs...)
	for _, i := range vs {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddView", i.GetViewId())
	}
	return v
}
func (v *FRadioGroup) Vertical() *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Orientation", "VERTICAL")
	return v
}
func (v *FRadioGroup) Horizontal() *FRadioGroup {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Orientation", "HORIZONTAL")
	return v
}
func (v *FRadioGroup) IsVertical() bool {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Orientation") == "VERTICAL"
}
func (v *FRadioGroup) OnCheckedChange(f func(int, *FRadioButton)) *FRadioGroup {
	fnId := NewToken()
	GlobalVars.EventHandlersMap[fnId] = func(vid string) string {
		for k, v := range v.Children {
			if v.VID == vid {
				f(k, v)
				break
			}
		}
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnCheckedChange", fnId)
	return v
}
