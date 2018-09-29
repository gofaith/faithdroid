package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		TextView(a).SetId("text").Size(-2, -1))
	Fab(a).Icon("drawable://add").BackgroundColor(Colors.Yellow).LayoutGravity(Gravitys.Bottom | Gravitys.Right).OnClick(func() {
		ShowToast(a, "clicked")
	}).Show()
}
