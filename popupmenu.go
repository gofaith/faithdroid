package faithdroid

type FPopupMenu struct {
	VID, ClassName string
	UI             string
}

func (base *FBaseView) popupMenu(menuItems ...interface{}) *FPopupMenu {
	v := &FPopupMenu{}
	v.VID = newToken()
	v.ClassName = "PopupMenu"
	v.UI = base.UI
	GlobalVars.uis[v.UI].NewView(v.ClassName, jsonArray([]string{v.VID, base.VID}))
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Menus", jsonArray(menuItems))
	return v
}
func (base *FBaseView) showPopupMenu(menuItems ...interface{}) {
	base.popupMenu(menuItems).show()
}
func (v *FPopupMenu) getVID() string {
	return v.VID
}
func (v *FPopupMenu) show() *FPopupMenu {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Show", "")
	return v
}
func (v *FPopupMenu) dismiss() *FPopupMenu {
	GlobalVars.uis[v.UI].ViewSetAttr(v.VID, "Dismiss", "")
	return v
}
