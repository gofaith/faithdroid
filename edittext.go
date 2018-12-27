package faithdroid

type FEditText struct {
	FBaseView
}

func EditText(a IActivity) *FEditText {
	v := &FEditText{}
	v.VID = NewToken()
	v.ClassName = "EditText"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetEditTextByItemId(id string) *FEditText {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FEditText); ok {
			return bt
		}
	}
	return nil
}

func (v *FEditText) Size(w, h int) *FEditText {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FEditText) X(x float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FEditText) Y(y float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FEditText) PivotX(x float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FEditText) PivotY(y float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FEditText) ScaleX(x float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FEditText) ScaleY(y float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FEditText) Rotation(r float64) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FEditText) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FEditText) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FEditText) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FEditText) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FEditText) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FEditText) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FEditText) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FEditText) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FEditText) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FEditText) SetId(s string) *FEditText {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FEditText) SetItemId(parent *FListView, id string) *FEditText {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetEditTextById(id string) *FEditText {
	if v, ok := GlobalVars.IdMap[id].(*FEditText); ok {
		return v
	}
	return nil
}

func (v *FEditText) Background(s string) *FEditText {
	v.FBaseView.Background(s)
	return v
}
func (v *FEditText) Foreground(s string) *FEditText {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FEditText) CachedForeground(s string) *FEditText {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FEditText) BackgroundColor(s string) *FEditText {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FEditText) CachedBackground(s string) *FEditText {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FEditText) Visible() *FEditText {
	v.FBaseView.Visible()
	return v
}
func (v *FEditText) Invisible() *FEditText {
	v.FBaseView.Invisible()
	return v
}
func (v *FEditText) Gone() *FEditText {
	v.FBaseView.Gone()
	return v
}

func (v *FEditText) Padding(left, top, right, bottom int) *FEditText {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FEditText) PaddingLeft(dp int) *FEditText {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FEditText) PaddingTop(dp int) *FEditText {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FEditText) PaddingRight(dp int) *FEditText {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FEditText) PaddingBottom(dp int) *FEditText {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FEditText) PaddingAll(dp int) *FEditText {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FEditText) Margin(left, top, right, bottom int) *FEditText {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FEditText) MarginLeft(dp int) *FEditText {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FEditText) MarginTop(dp int) *FEditText {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FEditText) MarginRight(dp int) *FEditText {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FEditText) MarginBottom(dp int) *FEditText {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FEditText) MarginAll(dp int) *FEditText {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FEditText) LayoutGravity(gravity int) *FEditText {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FEditText) Elevation(dp float64) *FEditText {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FEditText) Assign(fb **FEditText) *FEditText {
	*fb = v
	return v
}
func (v *FEditText) LayoutWeight(f int) *FEditText {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FEditText) OnTouch(f func(TouchEvent)) *FEditText {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FEditText) OnClick(f func()) *FEditText {
	v.FBaseView.OnClick(f)
	return v
}
func (v *FEditText) Clickable(b bool) *FEditText {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FEditText) Top2TopOf(id string) *FEditText {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FEditText) Top2BottomOf(id string) *FEditText {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FEditText) Bottom2TopOf(id string) *FEditText {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FEditText) Bottom2BottomOf(id string) *FEditText {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FEditText) Left2LeftOf(id string) *FEditText {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FEditText) Right2RightOf(id string) *FEditText {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FEditText) Left2RightOf(id string) *FEditText {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FEditText) Right2LeftOf(id string) *FEditText {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FEditText) CenterX() *FEditText {
	v.FBaseView.CenterX()
	return v
}
func (v *FEditText) CenterY() *FEditText {
	v.FBaseView.CenterY()
	return v
}
func (v *FEditText) WidthPercent(num float64) *FEditText {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FEditText) HeightPercent(num float64) *FEditText {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
func (v *FEditText) Text(s string) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FEditText) TextColor(s string) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}
func (v *FEditText) TextSize(dpsize int) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TextSize", SPrintf(dpsize))
	return v
}
func (v *FEditText) GetText() string {
	return GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FEditText) InputTypeSingleLineText() *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "InputType", "Text")
	return v
}
func (v *FEditText) InputTypeNumber() *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "InputType", "Number")
	return v
}
func (v *FEditText) InputTypePassword() *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "InputType", "Password")
	return v
}
func (v *FEditText) InputTypeEnglish() *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "InputType", "English")
	return v
}
func (v *FEditText) MaxLines(i int) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "MaxLines", SPrintf(i))
	return v
}
func (v *FEditText) MaxEms(i int) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "MaxEms", SPrintf(i))
	return v
}
func (v *FEditText) Hint(s string) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Hint", s)
	return v
}
func (v *FEditText) MaxLength(i int) *FEditText {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "MaxLength", SPrintf(i))
	return v
}
