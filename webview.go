package faithdroid

type FWebView struct {
	FBaseView
}

func WebView(a *Activity) *FWebView {
	v := &FWebView{}
	v.VID = NewToken()
	v.ClassName = "WebView"
	v.UI = a.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func WebViewItem(a *Activity, uri string) *FWebView {
	w := WebView(a)
	return w.Focusable(false).BackgroundColor(Colors.Transparent).Size(-1, -1).loadUri(uri)
}
func (vh *ViewHolder) GetWebViewByItemId(id string) *FWebView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FWebView); ok {
			return bt
		}
	}
	return nil
}

func (v *FWebView) Size(w, h int) *FWebView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FWebView) SetId(s string) *FWebView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FWebView) SetItemId(parent *FListView, id string) *FWebView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetWebViewById(id string) *FWebView {
	if v, ok := GlobalVars.IdMap[id].(*FWebView); ok {
		return v
	}
	return nil
}

func (v *FWebView) Background(s string) *FWebView {
	v.FBaseView.Background(s)
	return v
}
func (v *FWebView) BackgroundColor(s string) *FWebView {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FWebView) CachedBackground(s string) *FWebView {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FWebView) Visible() *FWebView {
	v.FBaseView.Visible()
	return v
}
func (v *FWebView) Invisible() *FWebView {
	v.FBaseView.Invisible()
	return v
}
func (v *FWebView) Gone() *FWebView {
	v.FBaseView.Gone()
	return v
}

func (v *FWebView) Padding(left, top, right, bottom int) *FWebView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FWebView) PaddingLeft(dp int) *FWebView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FWebView) PaddingTop(dp int) *FWebView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FWebView) PaddingRight(dp int) *FWebView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FWebView) PaddingBottom(dp int) *FWebView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FWebView) PaddingAll(dp int) *FWebView {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FWebView) Margin(left, top, right, bottom int) *FWebView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FWebView) MarginLeft(dp int) *FWebView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FWebView) MarginTop(dp int) *FWebView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FWebView) MarginRight(dp int) *FWebView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FWebView) MarginBottom(dp int) *FWebView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FWebView) MarginAll(dp int) *FWebView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FWebView) LayoutGravity(gravity int) *FWebView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FWebView) Elevation(dp float32) *FWebView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FWebView) Assign(fb **FWebView) *FWebView {
	*fb = v
	return v
}
func (v *FWebView) LayoutWeight(f int) *FWebView {
	v.FBaseView.LayoutWeight(f)
	return v
}

// --------------------------------------------------------

func (v *FWebView) loadUri(s string) *FWebView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Uri", s)
	return v
}
func (v *FWebView) Focusable(b bool) *FWebView {
	if b {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Focusable", "true")
	} else {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Focusable", "false")
	}
	return v
}
