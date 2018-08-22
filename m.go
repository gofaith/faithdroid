package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).deferShow().size(-2, -2).cachedBackground("https://www.baidu.com/img/bd_logo1.png").append(
		button(a).size(-2, -1).text("ahskjdg").backgroundColor("#0088ff").id("bt"),
		textview(a).size(-2, -1).text("click me").onClick(func() {
			getButtonById("bt").enabled(false)
		}))
}
