package faithdroid

func (a *MainActivity) OnCreate() {
	LinearLayout(a).DeferShow().Size(-2, -2).Append(
		Button(a).Text("permission").OnClick(func() {
			RequestPermissions(a, []string{Permissions.READ_EXTERNAL_STORAGE}, nil)
		}),
		Button(a).Text("open").OnClick(func() {
			OpenFile(a, "/storage/emulated/0/Download/app.apk")
		}),
		Button(a).Text("scan").OnClick(func() {
			ScanFile(a, "/storage/emulated/0/a.png")
		}))
}
