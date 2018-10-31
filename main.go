package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).Size(-2, -1).Text("get machine id").OnClick(func() {
			id := getUniqueID(a)
			AlertDialog(a).Append(TextView(a).Text(id)).Show()
		}))
}
