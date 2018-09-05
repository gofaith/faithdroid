package faithdroid

type FMenuItem struct {
	MyTitle   string
	MyIcon    string
	MyOnClick string
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
func (m *FMenuItem) onClick(f func()) *FMenuItem {
	fnId := newToken()
	m.MyOnClick = fnId
	GlobalVars.eventHandlersMap[fnId] = func(string) string {
		f()
		return ""
	}
	return m
}
func (m *FMenuItem) icon(s string) *FMenuItem {
	m.MyIcon = s
	return m
}
func SubMenu(title string, menuItems ...interface{}) *FSubMenu {
	m := &FSubMenu{}
	m.MyTitle = title
	m.MySubMenu = menuItems
	return m
}
