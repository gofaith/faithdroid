package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).background("assets://a1.jpg").append(
		button(a).size(-2, -1).text("ahskjdg").backgroundColor("#0088ff"))
}
