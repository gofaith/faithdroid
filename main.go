package faithdroid

type MainActivity struct {
	Activity
}

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	linearlayout(a).size(-2, -2).marginAll(100).backgroundColor(Colors.White).elevation(10).show()
}
