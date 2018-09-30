package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		Button(a).Size(-2, -1).Text("standard").OnClick(func() {
			StartActivity(a, func(a1 *Activity) {
				name := a1.GetExtraValue("name")
				TextView(a1).Text("standard activity by :" + name).Show()
			}, NewActivityConfig().LaunchMode_Standard().PutExtra("name", "steven"))
		}))
}
