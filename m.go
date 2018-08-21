package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().append(
		button(a).text("one").size(-1, -1).onClick(func() {
			getButtonById("s").text("changed")
		}),
		button(a).text("two").id("s").size(-2, -2))
}
