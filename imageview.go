package faithdroid

type FImageView struct {
	FBaseView
}

func ImageView(a *Activity) *FImageView {
	v := &FImageView{}
	v.VID = newToken()
	v.ClassName = "ImageView"
	v.UI = a.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.viewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) getImageViewByItemId(id string) *FImageView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.viewMap[v].(*FImageView); ok {
			return bt
		}
	}
	return nil
}

func (v *FImageView) size(w, h int) *FImageView {
	i := []int{w, h}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Size", jsonArray(i))
	return v
}
func (v *FImageView) setId(s string) *FImageView {
	GlobalVars.idMap[s] = v
	return v
}

func (v *FImageView) setItemId(parent *FListView, id string) *FImageView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.getVID()
	return v
}
func getImageViewById(id string) *FImageView {
	if v, ok := GlobalVars.idMap[id].(*FImageView); ok {
		return v
	}
	return nil
}

func (v *FImageView) background(s string) *FImageView {
	v.FBaseView.background(s)
	return v
}
func (v *FImageView) backgroundColor(s int) *FImageView {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FImageView) cachedBackground(s string) *FImageView {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FImageView) visible() *FImageView {
	v.FBaseView.visible()
	return v
}
func (v *FImageView) invisible() *FImageView {
	v.FBaseView.invisible()
	return v
}
func (v *FImageView) gone() *FImageView {
	v.FBaseView.gone()
	return v
}

func (v *FImageView) padding(left, top, right, bottom int) *FImageView {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FImageView) paddingLeft(dp int) *FImageView {
	v.FBaseView.padding(dp, 0, 0, 0)
	return v
}
func (v *FImageView) paddingTop(dp int) *FImageView {
	v.FBaseView.padding(0, dp, 0, 0)
	return v
}
func (v *FImageView) paddingRight(dp int) *FImageView {
	v.FBaseView.padding(0, 0, dp, 0)
	return v
}
func (v *FImageView) paddingBottom(dp int) *FImageView {
	v.FBaseView.padding(0, 0, 0, dp)
	return v
}
func (v *FImageView) paddingAll(dp int) *FImageView {
	v.FBaseView.padding(dp, dp, dp, dp)
	return v
}

func (v *FImageView) margin(left, top, right, bottom int) *FImageView {
	v.FBaseView.margin(left, top, right, bottom)
	return v
}
func (v *FImageView) marginLeft(dp int) *FImageView {
	v.FBaseView.margin(dp, 0, 0, 0)
	return v
}
func (v *FImageView) marginTop(dp int) *FImageView {
	v.FBaseView.margin(0, dp, 0, 0)
	return v
}
func (v *FImageView) marginRight(dp int) *FImageView {
	v.FBaseView.margin(0, 0, dp, 0)
	return v
}
func (v *FImageView) marginBottom(dp int) *FImageView {
	v.FBaseView.margin(0, 0, 0, dp)
	return v
}
func (v *FImageView) marginAll(dp int) *FImageView {
	v.FBaseView.margin(dp, dp, dp, dp)
	return v
}

func (v *FImageView) layoutGravity(gravity int) *FImageView {
	v.FBaseView.layoutGravity(gravity)
	return v
}
func (v *FImageView) elevation(dp float32) *FImageView {
	v.FBaseView.elevation(dp)
	return v
}

func (v *FImageView) assign(fb **FImageView) *FImageView {
	*fb = v
	return v
}
func (v *FImageView) layoutWeight(f int) *FImageView {
	v.FBaseView.layoutWeight(f)
	return v
}

// --------------------------------------------------------
func (v *FImageView) src(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go downloadNetFile(s, "/data/data/"+GlobalVars.uis[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := newToken()
			GlobalVars.eventHandlersMap[fnID] = func(string) string {
				GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.uis[v.UI].RunOnUIThread(fnID)
		})
		return v
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}
func (v *FImageView) cachedSrc(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go cacheNetFile(s, "/data/data/"+GlobalVars.uis[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := newToken()
			GlobalVars.eventHandlersMap[fnID] = func(string) string {
				GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.uis[v.UI].RunOnUIThread(fnID)
		})
		return v
	}
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}
