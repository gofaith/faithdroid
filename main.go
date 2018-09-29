package faithdroid

func (a *MainActivity) OnCreate() {
	var name string
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		RadioGroup(a).Vertical().OnCheckedChange(func(i int, r *FRadioButton) {
			name = r.GetText()
		}).Append(
			RadioButton(a).Text("asd"),
			RadioButton(a).Text("steven"),
			RadioButton(a).Text("yif")),
		Button(a).Text("show selected").Size(-2, -1).OnClick(func() {
			ShowToast(a, name)
		}))
}
