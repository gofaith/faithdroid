package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	c := Clipboard(a)
	c.OnChange(func() {
		GetTextViewById("txt").Text(c.GetClipData())
	})
	LinearLayout(a).DeferShow().Append(
		TextView(a).SetId("txt").Text(c.GetClipData()),
		Button(a).Text("set to a").OnClick(func() {
			c.ClipData("a")
		}))
}
