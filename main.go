package faithdroid

func (a *MainActivity) OnCreate() {
	var apps []App
	var vl *FListView
	VScrollView(a).Size(-2, -2).DeferShow().Append(
		TextView(a).Text("loading").SetId("load"),
		VListView(a, func(lv *FListView) IView {
			return LinearLayout(a).Horizontal().Size(-2, -1).Foreground(Colors.RippleEffect).Append(
				ImageView(a).Size(50, 50).SetItemId(lv, "icon"),
				TextView(a).Size(-1, -1).SetItemId(lv, "title"),
				Button(a).Size(-2, -1).Text("s"))
		}, func(vh *ViewHolder, pos int) {
			vh.GetImageViewByItemId("icon").Src(apps[pos].Icon)
			vh.GetTextViewByItemId("title").Text(apps[pos].Title)
		}, func() int {
			return len(apps)
		}).Assign(&vl),
	)
	go func() {
		apps = GetInstalledApps(a)
		RunOnUIThread(a, func() {
			vl.NotifyDataSetChanged()
			GetTextViewById("load").Text(JsonArray(apps))
		})
	}()
}
