package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		CheckBox(a).Text("check me").OnCheckedChangeListener(func(b bool) {
			ShowToast(a, SPrintf(b))
		}),
	)
}
