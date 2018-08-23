package faithdroid

import (
	"fmt"
)

type FButton struct {
	FBaseView
}

func (v *FButton) getVID() string {
	return v.vID
}

func button(a *Activity) *FButton {
	v := &FButton{}
	v.vID = newToken()
	v.className = "Button"
	v.ui = a.ui
	v.ui.NewView(v.className, v.vID)
	return v
}
func (v *FButton) size(w, h int) *FButton {
	i := []int{w, h}
	v.ui.ViewSetAttr(v.vID, "Size", jsonArray(i))
	return v
}
func (v *FButton) setId(s string) *FButton {
	idMap[s] = v
	return v
}
func getButtonById(id string) *FButton {
	if v, ok := idMap[id].(*FButton); ok {
		return v
	}
	return nil
}

func (v *FButton) background(s string) *FButton {
	v.FBaseView.background(s)
	return v
}
func (v *FButton) backgroundColor(s string) *FButton {
	v.FBaseView.backgroundColor(s)
	return v
}
func (v *FButton) cachedBackground(s string) *FButton {
	v.FBaseView.cachedBackground(s)
	return v
}

func (v *FButton) visible() *FButton {
	v.FBaseView.visible()
	return v
}
func (v *FButton) invisible() *FButton {
	v.FBaseView.invisible()
	return v
}
func (v *FButton) gone() *FButton {
	v.FBaseView.gone()
	return v
}

func (v *FButton) padding(left, top, right, bottom int) *FButton {
	v.FBaseView.padding(left, top, right, bottom)
	return v
}
func (v *FButton) paddingLeft(left int) *FButton {
	v.FBaseView.padding(left, 0, 0, 0)
	return v
}
func (v *FButton) paddingTop(top int) *FButton {
	v.FBaseView.padding(0, top, 0, 0)
	return v
}
func (v *FButton) paddingRight(right int) *FButton {
	v.FBaseView.padding(0, 0, right, 0)
	return v
}
func (v *FButton) paddingBottom(bottom int) *FButton {
	v.FBaseView.padding(0, 0, 0, bottom)
	return v
}
func (v *FButton) paddingAll(all int) *FButton {
	v.FBaseView.padding(all, all, all, all)
	return v
}

// --------------------------------------------------------
func (v *FButton) text(s string) *FButton {
	v.ui.ViewSetAttr(v.vID, "Text", s)
	return v
}
func (v *FButton) textColor(s string) *FButton {
	v.ui.ViewSetAttr(v.vID, "TextColor", s)
	return v
}
func (v *FButton) textSize(dpsize int) *FButton {
	v.ui.ViewSetAttr(v.vID, "TextSize", fmt.Sprintf("%v", dpsize))
	return v
}
func (v *FButton) getText() string {
	return v.ui.ViewGetAttr(v.vID, "Text")
}
func (v *FButton) enabled(b bool) *FButton {
	if b {
		v.ui.ViewSetAttr(v.vID, "Enabled", "true")
	} else {
		v.ui.ViewSetAttr(v.vID, "Enabled", "false")
	}
	return v
}
func (v *FButton) isEnabled() bool {
	if v.ui.ViewGetAttr(v.vID, "Enabled") == "true" {
		return true
	}
	return false
}

func (v *FButton) onClick(f func()) *FButton {
	fnID := newToken()
	eventHandlersMap[fnID] = func(string) {
		f()
	}
	v.ui.ViewSetAttr(v.vID, "OnClick", fnID)
	return v
}
