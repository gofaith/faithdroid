package faithdroid

type FSpace struct {
	FBaseView
}

func Space(a *Activity) *FSpace {
	v := &FSpace{}
	v.VID = newToken()
	v.ClassName = "Space"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}

func (vh *ViewHolder) getSpaceByItemId(id string) *FSpace {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FSpace); ok {
			return bt
		}
	}
	return nil
}

func (v *FSpace) size(w, h int) *FSpace {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FSpace) setId(s string) *FSpace {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FSpace) setItemId(parent *FListView, id string) *FSpace {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getSpaceById(id string) *FSpace {
	if v, ok := GlobalVars.idMap[id].(*FSpace); ok {
		return v
	}
	return nil
}

func (v *FSpace) background(s string) *FSpace {
	v.FBaseView.background(s)
	return v
}
func (v *FSpace) backgroundColor(s int) *FSpace {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FSpace) cachedBackground(s string) *FSpace {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FSpace) visible() *FSpace {
	v.FBaseView.visible()
	return v
}
func (v *FSpace) invisible() *FSpace {
	v.FBaseView.invisible()
	return v
}
func (v *FSpace) gone() *FSpace {
	v.FBaseView.gone()
	return v
}

func (v *FSpace) padding(left, top, right, bottom int) *FSpace {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FSpace) paddingLeft(dp int) *FSpace {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FSpace) paddingTop(dp int) *FSpace {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FSpace) paddingRight(dp int) *FSpace {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FSpace) paddingBottom(dp int) *FSpace {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FSpace) paddingAll(dp int) *FSpace {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FSpace) margin(left, top, right, bottom int) *FSpace {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FSpace) marginLeft(dp int) *FSpace {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FSpace) marginTop(dp int) *FSpace {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FSpace) marginRight(dp int) *FSpace {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FSpace) marginBottom(dp int) *FSpace {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FSpace) marginAll(dp int) *FSpace {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FSpace) layoutGravity(gravity int) *FSpace {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FSpace) elevation(dp float32) *FSpace {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FSpace) assign(fb **FSpace) *FSpace {
	*fb = v
	return v
}
func (v *FSpace) layoutWeight(f int) *FSpace {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
