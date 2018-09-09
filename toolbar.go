package faithdroid

type FToolbar struct {
	FBaseView
	MyMenu []interface{}
}

// --------------------------------------------------------------------------------
func Toolbar(a *Activity) *FToolbar {
	v := &FToolbar{}
	v.VID = NewToken()
	v.ClassName = "Toolbar"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FToolbar) Size(w, h int) *FToolbar {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FToolbar) SetId(s string) *FToolbar {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FToolbar) SetItemId(parent *FListView, id string) *FToolbar {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetToolbarById(id string) *FToolbar {
	if v, ok := GlobalVars.IdMap[id].(*FToolbar); ok {
		return v
	}
	return nil
}

func (v *FToolbar) Background(s string) *FToolbar {
	v.FBaseView.Background(s)
	return v
}
func (v *FToolbar) BackgroundColor(s string) *FToolbar {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FToolbar) CachedBackground(s string) *FToolbar {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FToolbar) Visible() *FToolbar {
	v.FBaseView.Visible()
	return v
}
func (v *FToolbar) Invisible() *FToolbar {
	v.FBaseView.Invisible()
	return v
}
func (v *FToolbar) Gone() *FToolbar {
	v.FBaseView.Gone()
	return v
}

func (v *FToolbar) Padding(left, top, right, bottom int) *FToolbar {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FToolbar) PaddingLeft(dp int) *FToolbar {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) PaddingTop(dp int) *FToolbar {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FToolbar) PaddingRight(dp int) *FToolbar {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FToolbar) PaddingBottom(dp int) *FToolbar {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FToolbar) PaddingAll(dp int) *FToolbar {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) Margin(left, top, right, bottom int) *FToolbar {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FToolbar) MarginLeft(dp int) *FToolbar {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FToolbar) MarginTop(dp int) *FToolbar {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FToolbar) MarginRight(dp int) *FToolbar {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FToolbar) MarginBottom(dp int) *FToolbar {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FToolbar) MarginAll(dp int) *FToolbar {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FToolbar) LayoutGravity(gravity int) *FToolbar {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FToolbar) Elevation(dp float32) *FToolbar {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FToolbar) Assign(fb **FToolbar) *FToolbar {
	*fb = v
	return v
}
func (v *FToolbar) LayoutWeight(f int) *FToolbar {
	v.FBaseView.LayoutWeight(f)
	return v
}

// ----------------------------------------------------------
func (v *FToolbar) Title(s string) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Title", s)
	return v
}
func (v *FToolbar) Subtitle(s string) *FToolbar {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "SubTitle", s)
	return v
}
func (v *FToolbar) Menus(mis ...interface{}) *FToolbar {
	v.MyMenu = mis
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Menu", JsonArray(v.MyMenu))
	return v
}
