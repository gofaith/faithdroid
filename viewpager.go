package faithdroid

import (
	"fmt"
)

type FViewPager struct {
	FBaseView
}

func (vh *ViewHolder) GetViewPagerByItemId(id string) *FViewPager {
	if v, ok := vh.Vlist[id]; ok {
		if bt, ok := GlobalVars.ViewMap[v].(*FViewPager); ok {
			return bt
		}
	}
	return nil
}
func (v *FViewPager) SetId(s string) *FViewPager {
	GlobalVars.IdMap[s] = v
	return v
}

func (v *FViewPager) SetItemId(parent *FListView, id string) *FViewPager {
	if parent.Vh.Vlist == nil {
		parent.Vh.Vlist = make(map[string]string)
	}
	parent.Vh.Vlist[id] = v.GetViewId()
	return v
}
func GetViewPagerById(id string) *FViewPager {
	if v, ok := GlobalVars.IdMap[id].(*FViewPager); ok {
		return v
	}
	return nil
}
func ViewPager(a IActivity) *FViewPager {
	v := &FViewPager{}
	v.VID = NewToken()
	v.ClassName = "ViewPager"
	v.UI = a.GetMyActivity().UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, v.VID)
	GlobalVars.ViewMap[v.VID] = v
	return v
}
func (v *FViewPager) Size(w, h int) *FViewPager {
	i := []int{w, h}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Size", JsonArray(i))
	return v
}

func (v *FViewPager) X(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "X", SPrintf(x))
	return v
}
func (v *FViewPager) Y(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Y", SPrintf(y))
	return v
}
func (v *FViewPager) PivotX(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotX", SPrintf(x))
	return v
}
func (v *FViewPager) PivotY(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "PivotY", SPrintf(y))
	return v
}
func (v *FViewPager) ScaleX(x float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleX", SPrintf(x))
	return v
}
func (v *FViewPager) ScaleY(y float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "ScaleY", SPrintf(y))
	return v
}
func (v *FViewPager) Rotation(r float64) *FViewPager {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Rotation", SPrintf(r))
	return v
}

func (v *FViewPager) GetX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "X")
	return a2f(x)
}
func (v *FViewPager) GetY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Y")
	return a2f(x)
}
func (v *FViewPager) GetWidth() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Width")
	return a2f(x)
}
func (v *FViewPager) GetHeight() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Height")
	return a2f(x)
}
func (v *FViewPager) GetPivotX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotX")
	return a2f(x)
}
func (v *FViewPager) GetPivotY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "PivotY")
	return a2f(x)
}
func (v *FViewPager) GetScaleX() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleX")
	return a2f(x)
}
func (v *FViewPager) GetScaleY() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "ScaleY")
	return a2f(x)
}
func (v *FViewPager) GetRotation() float64 {
	x := GlobalVars.UIs[v.UI].ViewGetAttr(v.VID, "Rotation")
	return a2f(x)
}
func (v *FViewPager) Background(s string) *FViewPager {
	v.FBaseView.Background(s)
	return v
}
func (v *FViewPager) Foreground(s string) *FViewPager {
	v.FBaseView.Foreground(s)
	return v
}
func (v *FViewPager) CachedForeground(s string) *FViewPager {
	v.FBaseView.CachedForeground(s)
	return v
}
func (v *FViewPager) BackgroundColor(s string) *FViewPager {
	v.FBaseView.BackgroundColor(s)
	return v
}

func (v *FViewPager) CachedBackground(s string) *FViewPager {
	v.FBaseView.CachedBackground(s)
	return v
}

func (v *FViewPager) Visible() *FViewPager {
	v.FBaseView.Visible()
	return v
}
func (v *FViewPager) Invisible() *FViewPager {
	v.FBaseView.Invisible()
	return v
}
func (v *FViewPager) Gone() *FViewPager {
	v.FBaseView.Gone()
	return v
}

func (v *FViewPager) Padding(left, top, right, bottom int) *FViewPager {
	v.FBaseView.Padding(left, top, right, bottom)
	return v
}
func (v *FViewPager) PaddingLeft(dp int) *FViewPager {
	v.FBaseView.Padding(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) PaddingTop(dp int) *FViewPager {
	v.FBaseView.Padding(0, dp, 0, 0)
	return v
}
func (v *FViewPager) PaddingRight(dp int) *FViewPager {
	v.FBaseView.Padding(0, 0, dp, 0)
	return v
}
func (v *FViewPager) PaddingBottom(dp int) *FViewPager {
	v.FBaseView.Padding(0, 0, 0, dp)
	return v
}
func (v *FViewPager) PaddingAll(all int) *FViewPager {
	v.FBaseView.Padding(all, all, all, all)
	return v
}

func (v *FViewPager) Margin(left, top, right, bottom int) *FViewPager {
	v.FBaseView.Margin(left, top, right, bottom)
	return v
}
func (v *FViewPager) MarginLeft(dp int) *FViewPager {
	v.FBaseView.Margin(dp, 0, 0, 0)
	return v
}
func (v *FViewPager) MarginTop(dp int) *FViewPager {
	v.FBaseView.Margin(0, dp, 0, 0)
	return v
}
func (v *FViewPager) MarginRight(dp int) *FViewPager {
	v.FBaseView.Margin(0, 0, dp, 0)
	return v
}
func (v *FViewPager) MarginBottom(dp int) *FViewPager {
	v.FBaseView.Margin(0, 0, 0, dp)
	return v
}
func (v *FViewPager) MarginAll(dp int) *FViewPager {
	v.FBaseView.Margin(dp, dp, dp, dp)
	return v
}

func (v *FViewPager) LayoutGravity(gravity int) *FViewPager {
	v.FBaseView.LayoutGravity(gravity)
	return v
}
func (v *FViewPager) Elevation(dp float64) *FViewPager {
	v.FBaseView.Elevation(dp)
	return v
}

func (v *FViewPager) Assign(fb **FViewPager) *FViewPager {
	*fb = v
	return v
}
func (v *FViewPager) LayoutWeight(f int) *FViewPager {
	v.FBaseView.LayoutWeight(f)
	return v
}

func (v *FViewPager) OnTouch(f func(TouchEvent)) *FViewPager {
	v.FBaseView.OnTouch(f)
	return v
}
func (v *FViewPager) OnClick(f func()) *FViewPager {
	v.FBaseView.OnClick(f)
	return v
}
func (v *FViewPager) Clickable(b bool) *FViewPager {
	v.FBaseView.Clickable(b)
	return v
}

//constraint
func (v *FViewPager) Top2TopOf(id string) *FViewPager {
	v.FBaseView.Top2TopOf(id)
	return v
}
func (v *FViewPager) Top2BottomOf(id string) *FViewPager {
	v.FBaseView.Top2BottomOf(id)
	return v
}
func (v *FViewPager) Bottom2TopOf(id string) *FViewPager {
	v.FBaseView.Bottom2TopOf(id)
	return v
}
func (v *FViewPager) Bottom2BottomOf(id string) *FViewPager {
	v.FBaseView.Bottom2BottomOf(id)
	return v
}

func (v *FViewPager) Left2LeftOf(id string) *FViewPager {
	v.FBaseView.Left2LeftOf(id)
	return v
}
func (v *FViewPager) Right2RightOf(id string) *FViewPager {
	v.FBaseView.Right2RightOf(id)
	return v
}

func (v *FViewPager) Left2RightOf(id string) *FViewPager {
	v.FBaseView.Left2RightOf(id)
	return v
}
func (v *FViewPager) Right2LeftOf(id string) *FViewPager {
	v.FBaseView.Right2LeftOf(id)
	return v
}

func (v *FViewPager) CenterX() *FViewPager {
	v.FBaseView.CenterX()
	return v
}
func (v *FViewPager) CenterY() *FViewPager {
	v.FBaseView.CenterY()
	return v
}
func (v *FViewPager) WidthPercent(num float64) *FViewPager {
	v.FBaseView.WidthPercent(num)
	return v
}
func (v *FViewPager) HeightPercent(num float64) *FViewPager {
	v.FBaseView.HeightPercent(num)
	return v
}
// --------------------------------------------------------
type FPage struct {
	VID string
}

func Page(createView func() IView) *FPage {
	p := &FPage{}
	p.VID = NewToken()
	GlobalVars.EventHandlersMap[p.VID] = func(string) string {
		return createView().GetViewId()
	}
	return p
}

func (v *FViewPager) Pages(ps ...*FPage) *FViewPager {
	if ps == nil {
		return v
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Pages", JsonArray(ps))
	return v
}
func (v *FViewPager) OnGetPage(getView func(pos int) IView, getCount func() int) *FViewPager {
	fnId := NewToken()
	fn:=func(s string) string {
		i, e := a2i(s)
		if e != nil {
			print("OnCreateView():", e.Error())
			return ""
		}
		return getView(i).GetViewId()
	}
	fn2:=func(string) string {
		return SPrintf(getCount())
	}
	fnId2 := NewToken()
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMap[fnId2] = fn2
	GlobalVars.EventHandlersMapLock.Unlock()

	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnCreateView", fnId)
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnGetCount", fnId2)
	return v
}
func (v *FViewPager) BindTabLayoutById(id string) *FViewPager {
	if iview, ok := GlobalVars.IdMap[id]; ok {
		GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "TabLayout", iview.GetViewId())
	}
	return v
}
func (v *FViewPager) CurrentItem(i int, soomth bool) *FViewPager {
	i2 := 0
	if soomth {
		i2 = 1
	}
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "CurrentItem", JsonArray([]int{i, i2}))
	return v
}
func (v *FViewPager) OnPageSelected(f func(i int)) *FViewPager {
	fnId := NewToken()
	fn:=func(s string) string {
		i, e := a2i(s)
		if e != nil {
			fmt.Println("OnPageSelected() a2i error:", e)
			return ""
		}
		f(i)
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "OnPageSelected", fnId)
	return v
}
