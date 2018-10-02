package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		ViewPager(a).Size(-2, -2).SetId("vp").LayoutWeight(1).Pages(
			Page(func() IView {
				return TextView(a).Text("page one")
			}),
			Page(func() IView {
				return TextView(a).Text("page two")
			})).OnPageSelected(func(i int) {
			GetBottomNavById("bn").SelectedItem(i)
		}),
		BottomNav(a).SetId("bn").Size(-2, -1).Menus(
			MenuItem("one").Icon(`drawable://add`).OnClick(func() {
				GetViewPagerById("vp").CurrentItem(0, true)
			}),
			MenuItem("two").Icon(`drawable://add`).OnClick(func() {
				GetViewPagerById("vp").CurrentItem(1, true)
			})))
}
