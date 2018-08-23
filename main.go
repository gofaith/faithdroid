package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).vertical().marginAll(100).backgroundColor(colors.Amber).paddingAll(20).append(
		button(a).text("asd"),
		button(a).text("two"))
}
