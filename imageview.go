package faithdroid

type FImageView struct {
	FBaseView
}

func ImageView(a IActivity) *FImageView {
	v := &FImageView{}
	v.VID = NewToken()
	v.ClassName = "ImageView"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (vh *ViewHolder) GetImageViewByItemId(id string) *FImageView {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FImageView); ok {
			return bt
		}
	}
	return nil
}

var ScaleTypes = struct {
	Center       string
	CenterCrop   string
	CenterInside string
	FitCenter    string
	FitStart     string
	FitEnd       string
	FitXY        string
	Matrix       string
}{
	"Center",
	"CenterCrop",
	"CenterInside",
	"FitCenter",
	"FitStart",
	"FitEnd",
	"FitXY",
	"Matrix",
}

func (v *FImageView) Size(w, h int) *FImageView {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}
func (v *FImageView) X(x float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FImageView) Y(y float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FImageView) PivotX(x float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FImageView) PivotY(y float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FImageView) ScaleX(x float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FImageView) ScaleY(y float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FImageView) Rotation(r float64) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FImageView) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FImageView) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FImageView) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FImageView) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FImageView) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FImageView) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FImageView) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FImageView) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FImageView) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FImageView) SetId(s string) *FImageView {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FImageView) SetItemId(parent *FListView, id string) *FImageView {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetImageViewById(id string) *FImageView {
	if v, ok := GlobalVars.IdMap[id].(*FImageView); ok {
		return v
	}
	return nil
}

func (v *FImageView) Background(s string) *FImageView {
	v.FBaseView.Background(s)
	return v
}
func (v *FImageView) Foreground(s string) *FImageView {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FImageView) CachedForeground(s string) *FImageView {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FImageView) BackgroundColor(s string) *FImageView {
	v.FBaseView.BackgroundColor(s)
	return v
}
func (v *FImageView) CachedBackground(s string) *FImageView {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FImageView) Visible() *FImageView {
	v.FBaseView.Visible()
	return v
}
func (v *FImageView) Invisible() *FImageView {
	v.FBaseView.Invisible()
	return v
}
func (v *FImageView) Gone() *FImageView {
	v.FBaseView.Gone()
	return v
}

func (v *FImageView) Padding(left, top, right, bottom int) *FImageView {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FImageView) PaddingLeft(dp int) *FImageView {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FImageView) PaddingTop(dp int) *FImageView {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FImageView) PaddingRight(dp int) *FImageView {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FImageView) PaddingBottom(dp int) *FImageView {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FImageView) PaddingAll(dp int) *FImageView {
	v.FBaseView.Padding(dp, dp, dp, dp)
	return v
}

func (v *FImageView) Margin(left, top, right, bottom int) *FImageView {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FImageView) MarginLeft(dp int) *FImageView {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FImageView) MarginTop(dp int) *FImageView {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FImageView) MarginRight(dp int) *FImageView {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FImageView) MarginBottom(dp int) *FImageView {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FImageView) MarginAll(dp int) *FImageView {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FImageView) LayoutGravity(gravity int) *FImageView {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FImageView) Elevation(dp float64) *FImageView {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FImageView) Assign(fb **FImageView) *FImageView {
	*fb = v
	return v
}
func (v *FImageView) LayoutWeight(f int) *FImageView {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FImageView) OnTouch(f func(TouchEvent)) *FImageView {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FImageView) OnClick(f func()) *FImageView {
	fnID := NewToken()
	GlobalVars.EventHandlersMap[fnID] = func(string) string {
		f()
		return ""
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnClick", fnID)
	return v
}
func (v *FImageView) Clickable(b bool) *FImageView {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FImageView) Top2TopOf(id string) *FImageView {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FImageView) Top2BottomOf(id string) *FImageView {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FImageView) Bottom2TopOf(id string) *FImageView {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FImageView) Bottom2BottomOf(id string) *FImageView {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FImageView) Left2LeftOf(id string) *FImageView {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FImageView) Right2RightOf(id string) *FImageView {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FImageView) Left2RightOf(id string) *FImageView {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FImageView) Right2LeftOf(id string) *FImageView {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FImageView) CenterX() *FImageView {
	v.FBaseView.CenterX()
	return v
}
func (v *FImageView) CenterY() *FImageView {
	v.FBaseView.CenterY()
	return v
}
// --------------------------------------------------------
func (v *FImageView) Src(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go DownloadNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/tmp/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunUIThread(fnID)
		})
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}
func (v *FImageView) CachedSrc(s string) *FImageView {
	if len(s) > len("https://") && s[:len("http")] == "http" {
		go CacheNetFile(s, "/data/data/"+GlobalVars.UIs[v.UI].GetPkg()+"/cacheDir/", func(f string) {
			fnID := NewToken()
			GlobalVars.EventHandlersMap[fnID] = func(string) string {
				GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", "file://"+f)
				return ""
			}
			GlobalVars.UIs[v.UI].RunUIThread(fnID)
		})
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Src", s)
	return v
}
func (v *FImageView) ScaleType(s string) *FImageView {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleType", s)
	return v
}
