package faithdroid

func (a *MainActivity) OnCreate() {
	ConstraintLayout(a).DeferShow().Size(-2, -2).Append(
		TextView(a).Text("nihao").CenterX().CenterY())
}
