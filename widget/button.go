package widget

import (
	"github.com/gofaith/faithdroid/interfaces"
	"github.com/gofaith/faithdroid/util"
	"github.com/gofaith/faithdroid/vars"
)

type FButton struct {
	FBaseView
}

func Button(w interfaces.IWindow) *FButton {
	v := &FButton{}
	v.Vid = util.NewNumToken()
	v.ClassName = vars.Button
	v.UI = w.GetUI()
	v.UI.New(v.ClassName, v.Vid)
	vars.SetView(v.Vid, v)
	return v
}

//common

func (v *FButton) Size(w, h int) {
	v.FBaseView.Size(w, h)
}

func (v *FButton) X(x float64) {
	v.FBaseView.X(x)
}
func (v *FButton) Y(y float64) {
	v.FBaseView.Y(y)
}
func (v *FButton) PivotX(x float64) {
	v.FBaseView.PivotX(x)
}
func (v *FButton) PivotY(y float64) {
	v.FBaseView.PivotY(y)
}
func (v *FButton) ScaleX(x float64) {
	v.FBaseView.ScaleX(x)
}
func (v *FButton) ScaleY(y float64) {
	v.FBaseView.ScaleY(y)
}
func (v *FButton) Rotation(r float64) {
	v.FBaseView.Rotation(r)
}

func (v *FButton) GetX() float64 {
	return v.FBaseView.GetX()
}
func (v *FButton) GetY() float64 {
	return v.FBaseView.GetY()
}
func (v *FButton) GetWidth() float64 {
	return v.FBaseView.GetWidth()
}
func (v *FButton) GetHeight() float64 {
	return v.FBaseView.GetHeight()
}
func (v *FButton) GetPivotX() float64 {
	return v.FBaseView.GetPivotX()
}
func (v *FButton) GetPivotY() float64 {
	return v.FBaseView.GetPivotY()
}
func (v *FButton) GetScaleX() float64 {
	return v.FBaseView.GetScaleX()
}
func (v *FButton) GetScaleY() float64 {
	return v.FBaseView.GetScaleY()
}
func (v *FButton) GetRotation() float64 {
	return v.FBaseView.GetRotation()
}

// attr

func (v *FButton) Background(s string) {
	v.FBaseView.Background(s)
}
func (v *FButton) Foreground(s string) {
	v.FBaseView.Foreground(s)
}
func (v *FButton) BackgroundColor(s string) {
	v.FBaseView.BackgroundColor(s)
}
func (v *FButton) Visible(b bool) {
	v.FBaseView.Visible(b)
}
func (v *FButton) Gone() {
	v.FBaseView.Gone()
}
func (v *FButton) IsVisible() bool {
	return v.FBaseView.IsVisible()
}
func (v *FButton) IsGone() bool {
	return v.FBaseView.IsGone()
}
func (v *FButton) Padding(left, top, right, bottom int) {
	v.FBaseView.Padding(left, top, right, bottom)
}
func (v *FButton) Margin(left, top, right, bottom int) {
	v.FBaseView.Margin(left, top, right, bottom)
}
func (v *FButton) LayoutGravity(gravity int) {
	v.FBaseView.LayoutGravity(gravity)
}
func (v *FButton) Elevation(dp float64) {
	v.FBaseView.Elevation(dp)
}
func (v *FButton) LayoutWeight(f int) {
	v.FBaseView.LayoutWeight(f)
}
func (v *FButton) Clickable(b bool) {
	v.FBaseView.Clickable(b)
}
func (f *FButton) OnClick(ff func()) {
	f.FBaseView.OnClick(ff)
}

//constraint

func (v *FButton) Top2TopOf(i interfaces.IView) {
	v.FBaseView.Top2TopOf(i)
}

func (v *FButton) Top2BottomOf(i interfaces.IView) {
	v.FBaseView.Top2BottomOf(i)
}

func (v *FButton) Bottom2TopOf(i interfaces.IView) {
	v.FBaseView.Bottom2TopOf(i)
}

func (v *FButton) Bottom2BottomOf(i interfaces.IView) {
	v.FBaseView.Bottom2BottomOf(i)
}

func (v *FButton) Left2LeftOf(i interfaces.IView) {
	v.FBaseView.Left2LeftOf(i)
}

func (v *FButton) Right2RightOf(i interfaces.IView) {
	v.FBaseView.Right2RightOf(i)
}

func (v *FButton) Left2RightOf(i interfaces.IView) {
	v.FBaseView.Left2RightOf(i)
}

func (v *FButton) Right2LeftOf(i interfaces.IView) {
	v.FBaseView.Right2LeftOf(i)
}
func (v *FButton) CenterX() {
	v.FBaseView.CenterX()
}
func (v *FButton) CenterY() {
	v.FBaseView.CenterY()
}
func (v *FButton) WidthPercent(num float64) {
	v.FBaseView.WidthPercent(num)
}
func (v *FButton) HeightPercent(num float64) {
	v.FBaseView.HeightPercent(num)
}
