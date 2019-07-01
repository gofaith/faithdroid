package widget

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofaith/faithdroid/interfaces"

	"github.com/StevenZack/tools/strToolkit"

	"github.com/gofaith/faithdroid/util"

	"github.com/gofaith/faithdroid/vars"

	"github.com/gofaith/faithdroid/base"
)

type FBaseView struct {
	base.FBase
	srcCachedBackground string
	srcCacheForeground  string
}

func (f *FBaseView) Show() {
	f.UI.Show(f.Vid)
}

func (v *FBaseView) Size(w, h int) {
	v.UI.SetAttr(v.Vid, vars.Size, strToolkit.JsonArray([]int{w, h}))
}

func (v *FBaseView) X(x float64) {
	v.UI.SetAttr(v.Vid, vars.X, strconv.FormatFloat(x, 'f', 6, 64))
}
func (v *FBaseView) Y(y float64) {
	v.UI.SetAttr(v.Vid, vars.Y, strconv.FormatFloat(y, 'f', 6, 64))
}
func (v *FBaseView) PivotX(x float64) {
	v.UI.SetAttr(v.Vid, vars.PivotX, strconv.FormatFloat(x, 'f', 6, 64))
}
func (v *FBaseView) PivotY(y float64) {
	v.UI.SetAttr(v.Vid, vars.PivotY, strconv.FormatFloat(y, 'f', 6, 64))
}
func (v *FBaseView) ScaleX(x float64) {
	v.UI.SetAttr(v.Vid, vars.ScaleX, strconv.FormatFloat(x, 'f', 6, 64))
}
func (v *FBaseView) ScaleY(y float64) {
	v.UI.SetAttr(v.Vid, vars.ScaleY, strconv.FormatFloat(y, 'f', 6, 64))
}
func (v *FBaseView) Rotation(r float64) {
	v.UI.SetAttr(v.Vid, vars.Rotation, strconv.FormatFloat(r, 'f', 6, 64))
}

func (v *FBaseView) GetX() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.X))
}
func (v *FBaseView) GetY() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.Y))
}
func (v *FBaseView) GetWidth() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.Width))
}
func (v *FBaseView) GetHeight() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.Height))
}
func (v *FBaseView) GetPivotX() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.PivotX))
}
func (v *FBaseView) GetPivotY() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.PivotY))
}
func (v *FBaseView) GetScaleX() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.ScaleX))
}
func (v *FBaseView) GetScaleY() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.ScaleY))
}
func (v *FBaseView) GetRotation() float64 {
	return util.Atof(v.UI.GetAttr(v.Vid, vars.Rotation))
}

// attr
func (v *FBaseView) Background(s string) {
	if strings.HasPrefix(s, "http") {
		go func() {
			f, e := util.CacheNetFile(s, v.UI.GetPkg())
			if e != nil {
				fmt.Println("download error :", e)
				return
			}
			if v.srcCachedBackground == f {
				return
			}
			base.RunOnUIThread(v.UI, func() {
				v.UI.SetAttr(v.Vid, vars.Background, f)
				v.srcCachedBackground = f
			})
		}()
		return
	}
	v.UI.SetAttr(v.Vid, vars.Background, s)
}
func (v *FBaseView) Foreground(s string) {
	if strings.HasPrefix(s, "http") {
		go func() {
			f, e := util.CacheNetFile(s, v.UI.GetPkg())
			if e != nil {
				fmt.Println(" error :", e)
				return
			}
			if v.srcCacheForeground == f {
				return
			}
			base.RunOnUIThread(v.UI, func() {
				v.UI.SetAttr(v.Vid, vars.Foreground, f)
				v.srcCacheForeground = f
			})
		}()
		return
	}
	v.UI.SetAttr(v.Vid, vars.Foreground, s)
}
func (v *FBaseView) BackgroundColor(s string) {
	v.UI.SetAttr(v.Vid, vars.BackgroundColor, s)
}
func (v *FBaseView) Visible(b bool) {
	v.UI.SetAttr(v.Vid, vars.Visibility, strconv.FormatBool(b))
}
func (v *FBaseView) Gone() {
	v.UI.SetAttr(v.Vid, vars.Visibility, "gone")
}
func (v *FBaseView) IsVisible() bool {
	return "true" == v.UI.GetAttr(v.Vid, vars.Visibility)
}
func (v *FBaseView) IsGone() bool {
	return "gone" == v.UI.GetAttr(v.Vid, vars.Visibility)
}
func (v *FBaseView) Padding(left, top, right, bottom int) {
	v.UI.SetAttr(v.Vid, vars.Padding, strToolkit.JsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) Margin(left, top, right, bottom int) {
	v.UI.SetAttr(v.Vid, vars.Margin, strToolkit.JsonArray([]int{left, top, right, bottom}))
}
func (v *FBaseView) LayoutGravity(gravity int) {
	v.UI.SetAttr(v.Vid, vars.LayoutGravity, strconv.Itoa(gravity))
}
func (v *FBaseView) Elevation(dp float64) {
	v.UI.SetAttr(v.Vid, vars.Elevation, strconv.FormatFloat(dp, 'f', 6, 64))
}
func (v *FBaseView) LayoutWeight(f int) {
	v.UI.SetAttr(v.Vid, vars.LayoutWeight, strconv.Itoa(f))
}
func (v *FBaseView) Clickable(b bool) {
	v.UI.SetAttr(v.Vid, vars.Clickable, strconv.FormatBool(b))
}
func (f *FBaseView) OnClick(ff func()) {
	fnID := util.NewNumToken()
	vars.SetFn(fnID, func(string) string {
		ff()
		return ""
	})
	f.UI.SetAttr(f.Vid, vars.OnClick, strconv.Itoa(fnID))
}

//constraint

func (v *FBaseView) Top2TopOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Top2TopOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Top2BottomOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Top2BottomOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Bottom2TopOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Bottom2TopOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Bottom2BottomOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Bottom2BottomOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Left2LeftOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Left2LeftOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Right2RightOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Right2RightOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Left2RightOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Left2RightOf, strconv.Itoa(i.GetVid()))
}

func (v *FBaseView) Right2LeftOf(i interfaces.IView) {
	v.UI.SetAttr(v.Vid, vars.Right2LeftOf, strconv.Itoa(i.GetVid()))
}
func (v *FBaseView) CenterX() {
	v.UI.SetAttr(v.Vid, vars.CenterX, "")
}
func (v *FBaseView) CenterY() {
	v.UI.SetAttr(v.Vid, vars.CenterY, "")
}
func (v *FBaseView) WidthPercent(num float64) {
	v.UI.SetAttr(v.Vid, vars.WidthPercent, strconv.FormatFloat(num, 'f', 6, 64))
}
func (v *FBaseView) HeightPercent(num float64) {
	v.UI.SetAttr(v.Vid, vars.HeightPercent, strconv.FormatFloat(num, 'f', 6, 64))
}
