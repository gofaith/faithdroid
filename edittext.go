package faithdroid

type FEditText struct {
	FBaseView
}

func EditText(a *Activity) *FEditText {
	v := &FEditText{}
	v.VID = newToken()
	v.ClassName = "EditText"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) getEditTextByItemId(id string) *FEditText {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FEditText); ok {
			return bt
		}
	}
	return nil
}

func (v *FEditText) size(w, h int) *FEditText {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FEditText) setId(s string) *FEditText {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FEditText) setItemId(parent *FListView, id string) *FEditText {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getEditTextById(id string) *FEditText {
	if v, ok := GlobalVars.idMap[id].(*FEditText); ok {
		return v
	}
	return nil
}

func (v *FEditText) background(s string) *FEditText {
	v.FBaseView.background(s)
	return v
}
func (v *FEditText) backgroundColor(s int) *FEditText {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FEditText) cachedBackground(s string) *FEditText {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FEditText) visible() *FEditText {
	v.FBaseView.visible()
	return v
}
func (v *FEditText) invisible() *FEditText {
	v.FBaseView.invisible()
	return v
}
func (v *FEditText) gone() *FEditText {
	v.FBaseView.gone()
	return v
}

func (v *FEditText) padding(left, top, right, bottom int) *FEditText {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FEditText) paddingLeft(dp int) *FEditText {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FEditText) paddingTop(dp int) *FEditText {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FEditText) paddingRight(dp int) *FEditText {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FEditText) paddingBottom(dp int) *FEditText {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FEditText) paddingAll(dp int) *FEditText {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FEditText) margin(left, top, right, bottom int) *FEditText {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FEditText) marginLeft(dp int) *FEditText {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FEditText) marginTop(dp int) *FEditText {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FEditText) marginRight(dp int) *FEditText {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FEditText) marginBottom(dp int) *FEditText {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FEditText) marginAll(dp int) *FEditText {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FEditText) layoutGravity(gravity int) *FEditText {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FEditText) elevation(dp float32) *FEditText {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FEditText) assign(fb **FEditText) *FEditText {
	*fb = v
	return v
}
func (v *FEditText) layoutWeight(f int) *FEditText {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FEditText) text(s string) *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Text", s)
	return v
}
func (v *FEditText) textColor(s string) *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextColor", s)
	return v
}
func (v *FEditText) textSize(dpsize int) *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "TextSize", sPrintf(dpsize))
	return v
}
func (v *FEditText) getText() string {
	return GlobalVars.uis[v.UI].ViewGetAttr(v.VID, "Text")
}
func (v *FEditText) typeText() *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "InputType", "Text")
	return v
}
func (v *FEditText) typeNumber() *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "InputType", "Number")
	return v
}
func (v *FEditText) typePassword() *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "InputType", "Password")
	return v
}
func (v *FEditText) typeEnglish() *FEditText {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "InputType", "English")
	return v
}
