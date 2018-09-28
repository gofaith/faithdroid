package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).SetId("ctn").Size(-2, -2).Append(
		Button(a).Text("add-1").Size(-1, -1).OnClick(func() {
			GetLinearLayoutById("ctn").AddViewAt(
				Button(a).Text("minus1").Size(-2, -1).LayoutGravity(Gravitys.Bottom),
				-1,
			)
		}),
		Button(a).Text("add-2").Size(-1, -1).OnClick(func() {
			GetLinearLayoutById("ctn").AddViewAt(
				Button(a).Text("minus2").Size(-2, -1).LayoutGravity(Gravitys.Bottom),
				-2,
			)
		}),
	).Show()
}
