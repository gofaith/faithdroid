package faithdroid

type FListView struct {
	FBaseView
}
type TypeOnBindData struct {
	Str      string
	Position int
}

func (v *FListView) getVID() string {
	return v.VID
}
func (v *FListView) setId(s string) *FListView {
	GlobalVars.idMap[s] = v
	return v
}
func getListViewById(id string) *FListView {
	if v, ok := GlobalVars.idMap[id].(*FListView); ok {
		return v
	}
	return nil
}
func vlistview(a *Activity, getView func() string, bindData func(string, int), getCount func() int) *FListView {
	v := &FListView{}
	v.VID = newToken()
	v.ClassName = "VListView"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	fnId1 := newToken()
	GlobalVars.eventHandlersMap[fnId1] = func(string) string {
		return getView()
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnGetView", fnId1)

	fnId2 := newToken()
	GlobalVars.eventHandlersMap[fnId2] = func(str string) string {
		obd := TypeOnBindData{}
		unJson(str, &obd)
		bindData(obd.Str, obd.Position)
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnBindData", fnId2)

	fnId3 := newToken()
	GlobalVars.eventHandlersMap[fnId3] = func(str string) string {
		return sPrintf(getCount())
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId3)
	return v
}
func hlistview(a *Activity, getView func() string, bindData func(string, int), getCount func() int) *FListView {
	v := &FListView{}
	v.VID = newToken()
	v.ClassName = "HListView"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	fnId1 := newToken()
	GlobalVars.eventHandlersMap[fnId1] = func(string) string {
		return getView()
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnGetView", fnId1)

	fnId2 := newToken()
	GlobalVars.eventHandlersMap[fnId2] = func(str string) string {
		obd := TypeOnBindData{}
		unJson(str, &obd)
		bindData(obd.Str, obd.Position)
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnBindData", fnId2)

	fnId3 := newToken()
	GlobalVars.eventHandlersMap[fnId3] = func(str string) string {
		return sPrintf(getCount())
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId3)
	return v
}
func (v *FListView) size(w, h int) *FListView {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}

func (v *FListView) background(s string) *FListView {
	v.FBaseView.background(s)
	return v
}
func (v *FListView) backgroundColor(s int) *FListView {
	v.FBaseView.backgroundColor(s)
	return v
}

func (v *FListView) cachedBackground(s string) *FListView {
	v.FBaseView.cachedBackground(s)
	return v
}
func (v *FListView) onClick(f func()) *FListView {
	fnID := newToken()
	GlobalVars.eventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FListView) visible() *FListView {
	v.FBaseView.visible()
	return v
}
func (v *FListView) invisible() *FListView {
	v.FBaseView.invisible()
	return v
}
func (v *FListView) gone() *FListView {
	v.FBaseView.gone()
	return v
}

func (v *FListView) padding(left, top, right, bottom int) *FListView {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FListView) paddingLeft(dp int) *FListView {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FListView) paddingTop(dp int) *FListView {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FListView) paddingRight(dp int) *FListView {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FListView) paddingBottom(dp int) *FListView {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FListView) paddingAll(all int) *FListView {
	v.FBaseView.padding(all, all, all, all)
	return v
}

func (v *FListView) margin(left, top, right, bottom int) *FListView {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FListView) marginLeft(dp int) *FListView {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FListView) marginTop(dp int) *FListView {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FListView) marginRight(dp int) *FListView {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FListView) marginBottom(dp int) *FListView {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FListView) marginAll(dp int) *FListView {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FListView) layoutGravity(gravity int) *FListView {
	v.FBaseView.layoutGravity(gravity)
	return v
}

func (v *FListView) elevation(dp float32) *FListView {
	v.FBaseView.elevation(dp)
	return v
}
func (v *FListView) assign(fb **FListView) *FListView {
	*fb = v
	return v
}

func (v *FListView) layoutWeight(f int) *FListView {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------------------
