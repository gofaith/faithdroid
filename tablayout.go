package faithdroid

type FTabLayout struct {
	FBaseView
}

func (vh *ViewHolder) GetTabLayoutByItemId(id string) *FTabLayout {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FTabLayout); ok {
			return bt
		}
	}
	return nil
}

func (v *FTabLayout) SetId(s string)*FTabLayout {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FTabLayout) SetItemId(parent *FListView, id string) *FTabLayout {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetVID()
	return v
}
func GetTabLayoutById(id string) *FTabLayout {
	if v, ok := GlobalVars.IdMap[id].(*FTabLayout); ok {
		return v
	}
	return nil
}
func TabLayout(a *Activity) *FTabLayout {
	v := &FTabLayout{}
	v.VID = NewToken()
	v.ClassName = "TabLayout"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FTabLayout) Size(w, h int) *FTabLayout {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FTabLayout) Background(s string) *FTabLayout {
	v.FBaseView.Background(s)
	return v
}
func (v *FTabLayout) BackgroundColor(s int) *FTabLayout {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FTabLayout) CachedBackground(s string) *FTabLayout {
	v.FBaseView.CachedBackground(s)
	return v
}
func (v *FTabLayout) onClick(f func()) *FTabLayout {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}

func (v *FTabLayout) Visible() *FTabLayout {
	v.FBaseView.Visible()
	return v
}
func (v *FTabLayout) Invisible() *FTabLayout {
	v.FBaseView.Invisible()
	return v
}
func (v *FTabLayout) Gone() *FTabLayout {
	v.FBaseView.Gone()
	return v
}

func (v *FTabLayout) Padding(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FTabLayout) PaddingLeft(dp int) *FTabLayout {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) PaddingTop(dp int) *FTabLayout {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) PaddingRight(dp int) *FTabLayout {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) PaddingBottom(dp int) *FTabLayout {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) PaddingAll(all int) *FTabLayout {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FTabLayout) Margin(left, top, right, bottom int) *FTabLayout {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FTabLayout) MarginLeft(dp int) *FTabLayout {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FTabLayout) MarginTop(dp int) *FTabLayout {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FTabLayout) MarginRight(dp int) *FTabLayout {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FTabLayout) MarginBottom(dp int) *FTabLayout {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FTabLayout) MarginAll(dp int) *FTabLayout {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FTabLayout) LayoutGravity(gravity int) *FTabLayout {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FTabLayout) Elevation(dp float32) *FTabLayout {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FTabLayout) Assign(fb **FTabLayout) *FTabLayout {
	*fb = v
	return v
}
func (v *FTabLayout) LayoutWeight(f int) *FTabLayout {
	v.FBaseView.LayoutWeight(f)
	return v
}

// ------------------------------------------------------------

type FTab struct {
	Text, Icon string
}

func Tab(text string, i ...string) *FTab {
	t := &FTab{}
	t.Text = text
	if len(i) > 0 {
		t.Icon = i[0]
	}
	return t
}

// ----------------------------------------------------------

func (v *FTabLayout) Tabs(ts ...*FTab) *FTabLayout {
	for _, t := range ts {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "AddTab", JsonObject(t))
	}
	return v
}
