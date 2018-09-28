package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).SetId("ctn").Size(-2, -2).Append(
		Button(a).Text("add").Size(-2, -1).OnClick(func() {
			GetLinearLayoutById("ctn").Append(
				Button(a).Text("two").Size(-2, -1),
			)
		}),
	).Show()
}
