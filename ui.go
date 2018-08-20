package faithdroid

type UIInterface interface {
	NewView(viewName string, vID string)
	ViewSetAttr(vID string, attr string, value string)
	ViewGetAttr(vID string, attr string) string
	ShowOnRootView(vID string)
}
type BaseView struct {
	vID, className string
}
