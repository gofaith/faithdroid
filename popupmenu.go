package faithdroid

type FPopupMenu struct {
	VID, ClassName string
	UI             string
}

func (base *FBaseView) PopupMenu(menuItems ...interface{}) *FPopupMenu {
	v := &FPopupMenu{}
	v.VID = NewToken()
	v.ClassName = "PopupMenu"
	v.UI = base.UI
	GlobalVars.UIs[v.UI].NewView(v.ClassName, JsonArray([]string{v.VID, base.VID}))
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Menus", JsonArray(menuItems))
	return v
}
func (base *FBaseView) ShowPopupMenu(menuItems ...interface{}) {
	base.PopupMenu(menuItems).Show()
}
func (v *FPopupMenu) GetVID() string {
	return v.VID
}
func (v *FPopupMenu) Show() *FPopupMenu {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Show", "")
	return v
}
func (v *FPopupMenu) Dismiss() *FPopupMenu {
	GlobalVars.UIs[v.UI].ViewSetAttr(v.VID, "Dismiss", "")
	return v
}
