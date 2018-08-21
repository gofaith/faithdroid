package faithdroid

type UIInterface interface {
	NewView(viewName string, vID string)
	ViewSetAttr(vID string, attr string, value string)
	ViewGetAttr(vID string, attr string) string
	ShowOnRootView(vID string)
}
type FBaseView struct {
	vID, className string
	ui             UIInterface
}
type IView interface {
	getVID() string
}

func (v *FBaseView) show() {
	v.ui.ShowOnRootView(v.vID)
}

var (
	idMap = make(map[string]IView)
)
