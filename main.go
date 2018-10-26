package faithdroid

func (a *MainActivity) OnCreate() {
	ConstraintLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).Text("top").Size(-2, -1).SetId("bt1").Top2TopOf(parent),
		Button(a).Text("second").Size(-1, -1).Top2BottomOf("bt1").SetId("bt2"),
		Button(a).Text("right").Size(-1, -1).Left2RightOf("bt2").SetId("bt3").Top2BottomOf("bt2"),
		Button(a).Text("center").Top2BottomOf("bt3").LayoutGravity(Gravitys.Center))
}
