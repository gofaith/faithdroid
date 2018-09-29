package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).Size(-2, -2).Append(
		TextView(a).Size(-2,-1).Text("hello").BackgroundColor()
	).Show()
}
