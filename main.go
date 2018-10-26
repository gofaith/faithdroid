package faithdroid

func (a *MainActivity) OnCreate() {
	ConstraintLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).BackgroundColor(Colors.White).Text("top").Size(-2, -1).SetId("bt1").Top2TopOf(parent),
		Button(a).BackgroundColor(Colors.White).Text("second").Size(-1, -1).Top2BottomOf("bt1").SetId("bt2"),
		Button(a).BackgroundColor(Colors.White).Text("right").Size(-1, -1).Left2RightOf("bt2").SetId("bt3").Top2BottomOf("bt2"),
		Button(a).BackgroundColor(Colors.White).Text("center").Top2BottomOf("bt3").CenterX(),
		Button(a).BackgroundColor(Colors.White).Text("centerY").Size(-2, 0).CenterY().HeightPercent(0.5))
}
