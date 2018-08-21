package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	var txt *FTextView
	txt = textview(a).text("txt").size(-2, -1)
	linearlayout(a).deferShow().append(
		button(a).text("one").size(-1, -1).onClick(func() {
			txt.text("hello changed")
		}),
		txt,
		button(a).text("two").size(-2, -2))
}
