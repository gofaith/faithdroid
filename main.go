package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		Toolbar(a).Size(-2, -1).Title("title").Menus(
			MenuItem("add").Icon("drawable://add").ShowAsAction().OnClick(func() {
				ShowToast(a, "search")
			}),
			MenuItem("chat"),
			MenuItem("me"),
			MenuItem("hello"),
			SubMenu("settings",
				MenuItem("set 1"),
				MenuItem("set 2"))))
}
