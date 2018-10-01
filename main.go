package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		TextView(a).SetId("text").Text("text view"),
		Button(a).Size(-2, -1).Text("change text").OnClick(func() {
			GetTextViewById("text").Text("text changed")
		}))
}
