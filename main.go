package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).vertical().marginAll(100).backgroundColor(Colors.Amber).paddingAll(20).append(
		button(a).size(-1, -1).text("s").layoutGravity(Gravitys.End),
		button(a).text("two"))
}
