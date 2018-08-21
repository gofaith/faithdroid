package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().append(
		button(a).text("one").size(-1, -1),
		textview(a).text("txt").size(-2, -1),
		button(a).text("two").size(-2, -2))
}
