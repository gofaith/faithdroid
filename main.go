package faithdroid

func (m *MainActivity) OnCreate() {
	a := &m.Activity
	LinearLayout(a).DeferShow().Append(
		TextView(a).SetId("txt"),
		Button(a).Text("check").OnClick(func() {
			b := CheckSelfPermission(a, Permissions.READ_PHONE_STATE)
			ShowToast(a, SPrintf(b))
		}),
		Button(a).Text("request").OnClick(func() {
			RequestPermissions(a,
				[]string{Permissions.CAMERA, Permissions.WRITE_EXTERNAL_STORAGE, Permissions.READ_PHONE_STATE},
				func(bs []bool) {
					ShowToast(a, JsonArray(bs))
				})
		}))
}
