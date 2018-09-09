package faithdroid

type FSpace struct {
	FBaseView
}

func Space(a *Activity) *FSpace {
	v := &FSpace{}
	v.VID = NewToken()
	v.ClassName = "Space"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}

func (vh *ViewHolder) GetSpaceByItemId(id string) *FSpace {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FSpace); ok {
			return bt
		}
	}
	return nil
}

func (v *FSpace) Size(w, h int) *FSpace {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FSpace) SetId(s string) *FSpace {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FSpace) SetItemId(parent *FListView, id string) *FSpace {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetSpaceById(id string) *FSpace {
	if v, ok := GlobalVars.IdMap[id].(*FSpace); ok {
		return v
	}
	return nil
}

func (v *FSpace) Background(s string) *FSpace {
	v.FBaseView.Background(s)
	return v
}
func (v *FSpace) BackgroundColor(s string) *FSpace {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FSpace) CachedBackground(s string) *FSpace {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FSpace) Visible() *FSpace {
	v.FBaseView.Visible()
	return v
}
func (v *FSpace) Invisible() *FSpace {
	v.FBaseView.Invisible()
	return v
}
func (v *FSpace) Gone() *FSpace {
	v.FBaseView.Gone()
	return v
}

func (v *FSpace) Padding(left, top, right, bottom int) *FSpace {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FSpace) PaddingLeft(dp int) *FSpace {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FSpace) PaddingTop(dp int) *FSpace {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FSpace) PaddingRight(dp int) *FSpace {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FSpace) PaddingBottom(dp int) *FSpace {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FSpace) PaddingAll(dp int) *FSpace {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FSpace) Margin(left, top, right, bottom int) *FSpace {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FSpace) MarginLeft(dp int) *FSpace {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FSpace) MarginTop(dp int) *FSpace {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FSpace) MarginRight(dp int) *FSpace {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FSpace) MarginBottom(dp int) *FSpace {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FSpace) MarginAll(dp int) *FSpace {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FSpace) LayoutGravity(gravity int) *FSpace {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FSpace) Elevation(dp float32) *FSpace {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FSpace) Assign(fb **FSpace) *FSpace {
	*fb = v
	return v
}
func (v *FSpace) LayoutWeight(f int) *FSpace {
	v.FBaseView.LayoutWeight(f)
	return v
}

// --------------------------------------------------------
