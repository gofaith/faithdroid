package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	var bt1, bt2, bt3 *FButton
	ConstraintLayout(a).Size(-2, -2).DeferShow().Append(
		Button(a).Assign(&bt1).Size(-1, -1).Text("one").LeftToLeft(Parent).TopToTop(Parent),
		Button(a).Assign(&bt2).Size(-1, -1).Text("two").LeftToRight(bt1).TopToTop(Parent).RightToRight(Parent),
		Button(a).Assign(&bt3).Size(0, 0).Text("three").LeftToRight(bt1).RightToRight(Parent).TopToBottom(bt2).BottomToBottom(Parent),
	)
}
