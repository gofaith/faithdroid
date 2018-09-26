package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).Append(
		Button(a).Text("copy").SetId("cp").OnClick(func() {
			Clipboard(a).ClipData("ok")
			ShowToast(a, "copied")
		}),
	).Show()
}
