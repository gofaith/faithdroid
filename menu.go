package faithdroid

type FMenuItem struct {
	MyTitle        string
	MyIcon         string
	MyOnClick      string
	MyShowAsAction string
}
type FSubMenu struct {
	MyTitle   string
	MySubMenu []interface{}
}

func MenuItem(title string) *FMenuItem {
	mi := &FMenuItem{}
	mi.MyTitle = title
	return mi
}
func (m *FMenuItem) OnClick(f func()) *FMenuItem {
	fnId := NewToken()
	m.MyOnClick = fnId
	fn:=func(string) string {
		f()
		return ""
	}
	GlobalVars.EventHandlersMapLock.Lock()
	GlobalVars.EventHandlersMap[fnId] = fn
	GlobalVars.EventHandlersMapLock.Unlock()
	return m
}
func (m *FMenuItem) Icon(s string) *FMenuItem {
	m.MyIcon = s
	return m
}
func (m *FMenuItem) ShowAsAction() *FMenuItem {
	m.MyShowAsAction = "IF_ROOM"
	return m
}
func SubMenu(title string, menuItems ...interface{}) *FSubMenu {
	m := &FSubMenu{}
	m.MyTitle = title
	m.MySubMenu = menuItems
	return m
}
