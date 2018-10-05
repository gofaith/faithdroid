package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		Button(a).Text("asd").SetId("bt").Size(-2, -1),
		Button(a).Text("toggle").OnClick(func() {
			bt := GetButtonById("bt")
			if bt.IsVisible() {
				bt.Invisible()
			} else {
				bt.Visible()
			}
		}))
}
