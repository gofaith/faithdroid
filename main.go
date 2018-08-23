package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).vertical().paddingAll(100).append(
		button(a).text("asd"),
		button(a).text("two"))
}
