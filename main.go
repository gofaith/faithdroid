package faithdroid

func (a *MainActivity) OnCreate() {
	ConstraintLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).BackgroundColor(Colors.Grey).Size(-2, -1).SetId("1st").Text("wide"),
		Button(a).BackgroundColor(Colors.Grey).SetId("2nd").Text("right").Top2BottomOf("1st").Right2RightOf(parent),
		Button(a).BackgroundColor(Colors.Grey).SetId("3rd").Text("centerX").CenterX().Top2BottomOf("2nd"),
		Button(a).BackgroundColor(Colors.Grey).SetId("4th").Text("centerY").CenterY().HeightPercent(0.5).WidthPercent(1.0),
		Button(a).BackgroundColor(Colors.Grey).SetId("half").Text("half width of parent").WidthPercent(0.5).Top2BottomOf("4th"))
}
