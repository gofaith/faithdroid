package faithdroid

type FToolbar struct {
	FBaseView
}

func toolbar(a *Activity) *FToolbar {
	v := &FToolbar{}
	v.VID = newToken()
	v.ClassName = "Toolbar"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (v *FToolbar) size(w, h int) *FToolbar {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FToolbar) setId(s string) *FToolbar {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FToolbar) setItemId(parent *FListView, id string) *FToolbar {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getToolbarById(id string) *FToolbar {
	if v, ok := GlobalVars.idMap[id].(*FToolbar); ok {
		return v
	}
	return nil
}

func (v *FToolbar) background(s string) *FToolbar {
	v.FBaseView.background(s)
	return v
}
func (v *FToolbar) backgroundColor(s int) *FToolbar {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FToolbar) cachedBackground(s string) *FToolbar {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FToolbar) visible() *FToolbar {
	v.FBaseView.visible()
	return v
}
func (v *FToolbar) invisible() *FToolbar {
	v.FBaseView.invisible()
	return v
}
func (v *FToolbar) gone() *FToolbar {
	v.FBaseView.gone()
	return v
}

func (v *FToolbar) padding(left, top, right, bottom int) *FToolbar {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FToolbar) paddingLeft(dp int) *FToolbar {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) paddingTop(dp int) *FToolbar {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FToolbar) paddingRight(dp int) *FToolbar {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FToolbar) paddingBottom(dp int) *FToolbar {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FToolbar) paddingAll(dp int) *FToolbar {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) margin(left, top, right, bottom int) *FToolbar {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FToolbar) marginLeft(dp int) *FToolbar {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) marginTop(dp int) *FToolbar {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FToolbar) marginRight(dp int) *FToolbar {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FToolbar) marginBottom(dp int) *FToolbar {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FToolbar) marginAll(dp int) *FToolbar {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) layoutGravity(gravity int) *FToolbar {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FToolbar) elevation(dp float32) *FToolbar {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FToolbar) assign(fb **FToolbar) *FToolbar {
	*fb = v
	return v
}
func (v *FToolbar) layoutWeight(f int) *FToolbar {
	v.FBaseView.layoutWeight(f)
	return v
}

// ----------------------------------------------------------
func (v *FToolbar) title(s string) *FToolbar {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FToolbar) subtitle(s string) *FToolbar {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "SubTitle", s)
	return v
}
