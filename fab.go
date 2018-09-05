package faithdroid

type FFab struct {
	FBaseView
}

func Fab(a *Activity) *FFab {
	v := &FFab{}
	v.VID = newToken()
	v.ClassName = "Fab"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) getFabByItemId(id string) *FFab {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FFab); ok {
			return bt
		}
	}
	return nil
}

func (v *FFab) size(w, h int) *FFab {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FFab) setId(s string) *FFab {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FFab) setItemId(parent *FListView, id string) *FFab {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getFabById(id string) *FFab {
	if v, ok := GlobalVars.idMap[id].(*FFab); ok {
		return v
	}
	return nil
}

func (v *FFab) background(s string) *FFab {
	v.FBaseView.background(s)
	return v
}
func (v *FFab) backgroundColor(s int) *FFab {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FFab) cachedBackground(s string) *FFab {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FFab) visible() *FFab {
	v.FBaseView.visible()
	return v
}
func (v *FFab) invisible() *FFab {
	v.FBaseView.invisible()
	return v
}
func (v *FFab) gone() *FFab {
	v.FBaseView.gone()
	return v
}

func (v *FFab) padding(left, top, right, bottom int) *FFab {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FFab) paddingLeft(dp int) *FFab {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FFab) paddingTop(dp int) *FFab {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FFab) paddingRight(dp int) *FFab {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FFab) paddingBottom(dp int) *FFab {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FFab) paddingAll(dp int) *FFab {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FFab) margin(left, top, right, bottom int) *FFab {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FFab) marginLeft(dp int) *FFab {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FFab) marginTop(dp int) *FFab {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FFab) marginRight(dp int) *FFab {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FFab) marginBottom(dp int) *FFab {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FFab) marginAll(dp int) *FFab {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FFab) layoutGravity(gravity int) *FFab {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FFab) elevation(dp float32) *FFab {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FFab) assign(fb **FFab) *FFab {
	*fb = v
	return v
}
func (v *FFab) layoutWeight(f int) *FFab {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------

func (v *FFab) icon(s string) *FFab {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Icon", s)
	return v
}
func (v *FFab) onClick(f func()) *FFab {
	fnId := newToken()
	GlobalVars.eventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnId)
	return v
}
