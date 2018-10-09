package faithdroid

func (a *MainActivity) OnCreate() {
	s := `https://pp.myapp.com/ma_icon/0/icon_10966186_1538016476/256`
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		Toolbar(a).NavigationIcon(s).OnNavigationIconClick(func() {
			ShowToast(a, "clicked")
		}),
		Button(a).Text("open").OnClick(func() {
			StartActivity(a, createView, nil)
		}),
		ImageView(a).Src(s),
		Button(a).Size(-2, -1).Background(s).Foreground(Colors.RippleEffect))
}
func createView(a IActivity) {
	LinearLayout(a).Size(-2, -2).DeferShow().Append(
		Toolbar(a).Title("asd"))
}
