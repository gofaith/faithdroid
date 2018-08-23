package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).vertical().append(
		textview(a).text("asdkjh").setId("txt"),
		button(a).setId("bt").text("aksgdg").onClick(func() {
			getTextViewById("txt").textColor(colors.TealDark)
			getButtonById("bt").textColor(colors.White)
		}))
}
